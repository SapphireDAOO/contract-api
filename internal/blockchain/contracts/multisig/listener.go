package multisig

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/SapphireDAOO/contract-api/internal/discord"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

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

func (c *Multisig) subscribeLogs(ctx context.Context, query ethereum.FilterQuery,
	logs chan types.Log) ethereum.Subscription {
	return c.client.SubscribeLogs(ctx, query, logs, "Multisig")
}
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
			go c.notifier.SendEmbed(*embed)
		}
	}
}

func (c *Multisig) buildEmbed(ctx context.Context, vLog *types.Log) (*discord.Embed, error) {
	if len(vLog.Topics) == 0 {
		return nil, nil
	}

	base := discord.Embed{
		URL:    c.link("/tx/", vLog.TxHash.Hex()),
		Footer: &discord.Footer{Text: "Multisig " + discord.ShortHex(c.address.Hex()) + " • Base Sepolia"},
	}
	var lines []string

	switch vLog.Topics[0] {
	case transactionProposedTopic:
		event, err := c.contract.UnpackTransactionProposedEvent(vLog)
		if err != nil {
			return nil, err
		}
		act := c.decodeAction(event.Target, event.Data)
		state := c.lookupProposal(ctx, event.TxHash)
		base.Title = "📝 Proposal: " + actionTitle(act)
		base.Color = discord.ColorBlue
		lines = append(lines, progressLine(stageProposed, state.approvals, state.threshold))
		lines = append(lines,
			fmt.Sprintf("%s proposed %s on %s.",
				c.addressLink(event.Proposer), actionName(act), c.targetName(act, event.Target)))
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
		state := c.lookupProposal(ctx, event.TxHash)
		act := state.act

		// The event's own count is authoritative for this approval; the
		// contract read only supplies the threshold.
		approvals := event.ApprovalCount
		if approvals == nil {
			approvals = state.approvals
		}
		reached := thresholdReached(approvals, state.threshold)

		base.Color = discord.ColorYellow
		if reached {
			base.Title = "✅ Fully approved: " + actionTitle(act)
		} else {
			base.Title = "✍️ Approved by a signer: " + actionTitle(act)
		}

		lines = append(lines, progressLine(stageApprovals, approvals, state.threshold))

		approval := fmt.Sprintf("%s approved %s — **%s** so far.",
			c.addressLink(event.Approver), actionName(act), plural(approvals, "approval"))
		if reached {
			// The approval that reaches the threshold says so here rather
			// than in a second notification for TransactionApproved.
			approval += " That meets the threshold, so it is **ready to execute**."
		}
		lines = append(lines, approval, proposalLine(event.TxHash))

	case transactionApprovedTopic:
		// Reported by the approval that reached the threshold, so this event
		// would only duplicate it.
		return nil, nil

	case transactionExecutedTopic:
		event, err := c.contract.UnpackTransactionExecutedEvent(vLog)
		if err != nil {
			return nil, err
		}
		state := c.lookupProposal(ctx, event.TxHash)
		act := state.act
		base.Title = "🚀 Executed: " + actionTitle(act)
		base.Color = discord.ColorGreen
		lines = append(lines, progressLine(stageExecuted, state.approvals, state.threshold))
		lines = append(lines,
			fmt.Sprintf("%s executed %s. The change is now live on-chain.",
				c.addressLink(event.Executor), actionName(act)))
		lines = append(lines, actionArgLines(act)...)
		lines = append(lines, proposalLine(event.TxHash))

	case transactionCanceledTopic:
		event, err := c.contract.UnpackTransactionCanceledEvent(vLog)
		if err != nil {
			return nil, err
		}
		state := c.lookupProposal(ctx, event.TxHash)
		act := state.act
		base.Title = "🚫 Canceled: " + actionTitle(act)
		base.Color = discord.ColorRed
		lines = append(lines, progressLine(stageCanceled, state.approvals, state.threshold))
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
			fmt.Sprintf("%s is now a signer on the multisig.", c.addressLink(event.Signer)))

	case signerRemovedTopic:
		event, err := c.contract.UnpackSignerRemovedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "➖ Signer removed"
		base.Color = discord.ColorPurple
		lines = append(lines,
			fmt.Sprintf("%s was removed as a signer on the multisig.", c.addressLink(event.Signer)))

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

	lines = append(lines, fmt.Sprintf("[Open Multisig](%s) • [View on Basescan](%s)", c.dashboardURL, base.URL))
	// Blank lines between sections keep the message easy to scan; decoded
	// parameter bullets (actionArgLines) stay grouped as one block.
	base.Description = strings.Join(lines, "\n\n")
	return &base, nil
}

// proposal is what a lifecycle notification needs: what the transaction does,
// how many approvals it has, and how many it needs. Any field may be nil when
// the chain call fails, and the renderers degrade rather than error.
type proposal struct {
	act       *action
	approvals *big.Int
	threshold *big.Int
}

// lookupProposal reads a proposal from the multisig by its internal id.
func (c *Multisig) lookupProposal(ctx context.Context, txHash [32]byte) proposal {
	var p proposal

	if c.client == nil || c.client.HTTP == nil || c.instance == nil {
		return p
	}

	callCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts := &bind.CallOpts{Context: callCtx}

	tx, err := bind.Call(c.instance, opts, c.contract.PackGetTransaction(txHash), c.contract.UnpackGetTransaction)
	if err != nil {
		log.Printf("Failed to look up multisig transaction %s: %v", common.Hash(txHash).Hex(), err)
	} else {
		p.act = c.decodeAction(tx.Target, tx.Data)
		p.approvals = tx.ApprovalCount
	}

	threshold, err := bind.Call(c.instance, opts, c.contract.PackGetThreshold(), c.contract.UnpackGetThreshold)
	if err != nil {
		log.Printf("Failed to read the multisig threshold: %v", err)
	} else {
		p.threshold = threshold
	}

	return p
}

// stage is where in its lifecycle a proposal is, as told by the event being
// reported rather than by the contract's own status enum.
type stage int

const (
	stageProposed stage = iota
	stageApprovals
	stageExecuted
	stageCanceled
)

// progressLine renders the lifecycle as propose -> approvals -> execution,
// emphasising the step this notification is about, so a reader can see at a
// glance where the proposal stands.
func progressLine(current stage, approvals, threshold *big.Int) string {
	final := "🚀 Execution"
	switch current {
	case stageExecuted:
		final = "🚀 Executed"
	case stageCanceled:
		final = "🚫 Canceled"
	}

	steps := []string{"📝 Proposed", "✍️ " + approvalsText(approvals, threshold), final}

	at := 2
	switch current {
	case stageProposed:
		at = 0
	case stageApprovals:
		at = 1
	}
	steps[at] = "**" + steps[at] + "**"

	return strings.Join(steps, "  →  ")
}

func approvalsText(approvals, threshold *big.Int) string {
	switch {
	case approvals != nil && threshold != nil:
		return fmt.Sprintf("Approvals %s/%s", approvals, threshold)
	case approvals != nil:
		return fmt.Sprintf("Approvals %s", approvals)
	default:
		return "Approvals"
	}
}

// thresholdReached reports whether a proposal has the approvals it needs.
func thresholdReached(approvals, threshold *big.Int) bool {
	if approvals == nil || threshold == nil || threshold.Sign() <= 0 {
		return false
	}
	return approvals.Cmp(threshold) >= 0
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
func (c *Multisig) targetName(act *action, target common.Address) string {
	for _, kc := range c.known {
		if kc.address == target {
			return fmt.Sprintf("the %s (%s)", kc.name, c.addressLink(target))
		}
	}
	if act != nil {
		return fmt.Sprintf("the %s (%s)", act.Contract, c.addressLink(target))
	}
	return c.addressLink(target)
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
	return fmt.Sprintf("Transaction id: `%s`", discord.ShortHex(common.Hash(txHash).Hex()))
}

func (c *Multisig) addressLink(addr common.Address) string {
	return discord.AddressLink(c.explorerURL, addr)
}
func (c *Multisig) link(path, value string) string {
	return config.Link(c.explorerURL, path, value)
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
	if n == nil {
		return "0 " + unit + "s"
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return "1 " + unit
	}
	return n.String() + " " + unit + "s"
}
