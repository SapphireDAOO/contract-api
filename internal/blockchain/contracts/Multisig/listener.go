package multisig

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/orgs/SapphireDAOO/contract-api/internal/discord"
)

const explorerURL = "https://sepolia.basescan.org"
const dashboardURL = "https://sapphire-dao-website-git-feat-metric-data-bh-ead.vercel.app/multisig/"

var (
	transactionProposedTopic = crypto.Keccak256Hash([]byte("TransactionProposed(bytes32,address,uint256,bytes,uint256,address)"))
	approvalAddedTopic       = crypto.Keccak256Hash([]byte("ApprovalAdded(bytes32,address,uint256)"))
	transactionApprovedTopic = crypto.Keccak256Hash([]byte("TransactionApproved(bytes32)"))
	transactionExecutedTopic = crypto.Keccak256Hash([]byte("TransactionExecuted(bytes32,address)"))
	transactionCanceledTopic = crypto.Keccak256Hash([]byte("TransactionCanceled(bytes32)"))
	signerAddedTopic         = crypto.Keccak256Hash([]byte("SignerAdded(address)"))
	signerRemovedTopic       = crypto.Keccak256Hash([]byte("SignerRemoved(address)"))
	thresholdUpdatedTopic    = crypto.Keccak256Hash([]byte("ThresholdUpdated(uint256,uint256)"))
)

func (c *Multisig) subscribeLogs(ctx context.Context, query ethereum.FilterQuery, logs chan types.Log) ethereum.Subscription {
	for {
		sub, err := c.client.WS.SubscribeFilterLogs(ctx, query, logs)
		if err == nil {
			return sub
		}
		log.Printf("Failed to subscribe to Multisig logs: %v", err)

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(5 * time.Second):
		}
	}
}

// ListenToEvents subscribes to every log emitted by the multisig contract and
// posts a Discord notification for each one.
func (c *Multisig) ListenToEvents(ctx context.Context) {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		log.Printf("multisig listener disabled: client or contract address not initialized")
		return
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{*c.address},
	}

	logs := make(chan types.Log)
	sub := c.subscribeLogs(ctx, query, logs)
	if sub == nil {
		return
	}
	defer sub.Unsubscribe()

	log.Println("Listening for Multisig events...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Multisig listener stopping")
			return

		case err := <-sub.Err():
			log.Printf("Multisig subscription error: %v", err)
			sub.Unsubscribe()
			sub = c.subscribeLogs(ctx, query, logs)
			if sub == nil {
				return
			}

		case vLog := <-logs:
			embed, err := c.buildEmbed(ctx, &vLog)
			if err != nil {
				log.Printf("Failed to parse Multisig event: %v", err)
				continue
			}
			if embed == nil {
				continue
			}

			log.Printf("Multisig event: %s (%s)", embed.Title, vLog.TxHash.Hex())
			go discord.SendEmbed(*embed)
		}
	}
}

func (c *Multisig) buildEmbed(ctx context.Context, vLog *types.Log) (*discord.Embed, error) {
	if len(vLog.Topics) == 0 {
		return nil, nil
	}

	base := discord.Embed{
		URL:    explorerURL + "/tx/" + vLog.TxHash.Hex(),
		Footer: &discord.Footer{Text: "Multisig " + shortHex(c.address.Hex()) + " • Base Sepolia"},
	}
	var lines []string

	switch vLog.Topics[0] {
	case transactionProposedTopic:
		event, err := c.contract.UnpackTransactionProposedEvent(vLog)
		if err != nil {
			return nil, err
		}
		act := decodeAction(event.Target, event.Data)
		base.Title = "📝 Proposal: " + actionTitle(act)
		base.Color = discord.ColorBlue
		lines = append(lines,
			fmt.Sprintf("%s proposed %s on %s.",
				addressLink(event.Proposer), actionName(act), targetName(act, event.Target)))
		lines = append(lines, actionArgLines(act)...)
		if event.Value != nil && event.Value.Sign() > 0 {
			lines = append(lines, fmt.Sprintf("It sends **%s** along with the call.", formatEth(event.Value)))
		}
		lines = append(lines, proposalLine(event.TxHash)+fmt.Sprintf(" (nonce %s)", event.Nonce))

	case approvalAddedTopic:
		event, err := c.contract.UnpackApprovalAddedEvent(vLog)
		if err != nil {
			return nil, err
		}
		act := c.lookupAction(ctx, event.TxHash)
		base.Title = "✍️ Approved by a signer: " + actionTitle(act)
		base.Color = discord.ColorYellow
		lines = append(lines,
			fmt.Sprintf("%s approved %s — **%s** so far.",
				addressLink(event.Approver), actionName(act), plural(event.ApprovalCount, "approval")),
			proposalLine(event.TxHash),
		)

	case transactionApprovedTopic:
		event, err := c.contract.UnpackTransactionApprovedEvent(vLog)
		if err != nil {
			return nil, err
		}
		act := c.lookupAction(ctx, event.TxHash)
		base.Title = "✅ Fully approved: " + actionTitle(act)
		base.Color = discord.ColorYellow
		lines = append(lines,
			fmt.Sprintf("%s has enough approvals and is **ready to execute**.",
				capitalize(actionName(act))),
			proposalLine(event.TxHash),
		)

	case transactionExecutedTopic:
		event, err := c.contract.UnpackTransactionExecutedEvent(vLog)
		if err != nil {
			return nil, err
		}
		act := c.lookupAction(ctx, event.TxHash)
		base.Title = "🚀 Executed: " + actionTitle(act)
		base.Color = discord.ColorGreen
		lines = append(lines,
			fmt.Sprintf("%s executed %s. The change is now live on-chain.",
				addressLink(event.Executor), actionName(act)))
		lines = append(lines, actionArgLines(act)...)
		lines = append(lines, proposalLine(event.TxHash))

	case transactionCanceledTopic:
		event, err := c.contract.UnpackTransactionCanceledEvent(vLog)
		if err != nil {
			return nil, err
		}
		act := c.lookupAction(ctx, event.TxHash)
		base.Title = "🚫 Canceled: " + actionTitle(act)
		base.Color = discord.ColorRed
		lines = append(lines,
			fmt.Sprintf("%s was canceled and can no longer be executed.",
				capitalize(actionName(act))),
			proposalLine(event.TxHash),
		)

	case signerAddedTopic:
		event, err := c.contract.UnpackSignerAddedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "➕ Signer added"
		base.Color = discord.ColorPurple
		lines = append(lines,
			fmt.Sprintf("%s is now a signer on the multisig.", addressLink(event.Signer)))

	case signerRemovedTopic:
		event, err := c.contract.UnpackSignerRemovedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "➖ Signer removed"
		base.Color = discord.ColorPurple
		lines = append(lines,
			fmt.Sprintf("%s was removed as a signer on the multisig.", addressLink(event.Signer)))

	case thresholdUpdatedTopic:
		event, err := c.contract.UnpackThresholdUpdatedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "🔧 Approval threshold changed"
		base.Color = discord.ColorPurple
		lines = append(lines,
			fmt.Sprintf("The number of approvals required went from **%s** to **%s**.",
				event.OldThreshold, event.NewThreshold))

	default:
		return nil, nil
	}

	lines = append(lines, fmt.Sprintf("[Open Multisig](%s) • [View on Basescan](%s)", dashboardURL, base.URL))
	// Blank lines between sections keep the message easy to scan; decoded
	// parameter bullets (actionArgLines) stay grouped as one block.
	base.Description = strings.Join(lines, "\n\n")
	return &base, nil
}

// lookupAction fetches a proposal from the multisig by its internal id and
// decodes what it does, so approve/execute/cancel notifications can be named.
func (c *Multisig) lookupAction(ctx context.Context, txHash [32]byte) *action {
	if c.client == nil || c.client.HTTP == nil || c.instance == nil {
		return nil
	}

	callCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	data := c.contract.PackGetTransaction(txHash)
	tx, err := bind.Call(c.instance, &bind.CallOpts{Context: callCtx}, data, c.contract.UnpackGetTransaction)
	if err != nil {
		log.Printf("Failed to look up multisig transaction %s: %v", common.Hash(txHash).Hex(), err)
		return nil
	}
	return decodeAction(tx.Target, tx.Data)
}

// actionTitle names a decoded action for embed titles, e.g. "Set Fee Rate".
func actionTitle(act *action) string {
	if act == nil {
		return "Transaction"
	}
	return act.Name
}

// actionName names a decoded action mid-sentence, e.g. "**Set Fee Rate**".
func actionName(act *action) string {
	if act == nil {
		return "the transaction"
	}
	return "**" + act.Name + "**"
}

// targetName describes the contract a proposal calls into.
func targetName(act *action, target common.Address) string {
	for _, kc := range knownContracts {
		if kc.address == target {
			return fmt.Sprintf("the %s (%s)", kc.name, addressLink(target))
		}
	}
	if act != nil {
		return fmt.Sprintf("the %s (%s)", act.Contract, addressLink(target))
	}
	return addressLink(target)
}

// actionArgLines renders the decoded parameters as a single block,
// e.g. "• New Fee Rate: 300".
func actionArgLines(act *action) []string {
	if act == nil || len(act.Args) == 0 {
		return nil
	}
	lines := make([]string, 0, len(act.Args))
	for _, arg := range act.Args {
		lines = append(lines, fmt.Sprintf("• %s: **%s**", arg.Name, arg.Value))
	}
	return []string{strings.Join(lines, "\n")}
}

// proposalLine renders the multisig's internal transaction id, which groups
// the propose/approve/execute notifications for the same transaction.
func proposalLine(txHash [32]byte) string {
	return fmt.Sprintf("Transaction id: `%s`", shortHex(common.Hash(txHash).Hex()))
}

func addressLink(addr common.Address) string {
	return fmt.Sprintf("[`%s`](%s/address/%s)", shortHex(addr.Hex()), explorerURL, addr.Hex())
}

// shortHex shortens a hex string to the 0x1234…abcd form.
func shortHex(s string) string {
	if len(s) <= 12 {
		return s
	}
	return s[:6] + "…" + s[len(s)-4:]
}

func formatEth(wei *big.Int) string {
	if wei == nil || wei.Sign() == 0 {
		return "0 ETH"
	}
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return eth.Text('f', -1) + " ETH"
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func plural(n *big.Int, unit string) string {
	if n != nil && n.Cmp(big.NewInt(1)) == 0 {
		return "1 " + unit
	}
	return n.String() + " " + unit + "s"
}
