package multisig

import (
	"bytes"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"unicode"

	"github.com/SapphireDAOO/contract-api/internal/blockchain/gen/intermediatedpaymentprocessor"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/multisig"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/gen/paymentprocessorstorage"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/gen/simplepaymentprocessor"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type knownContract struct {
	name    string
	address common.Address
	abi     *abi.ABI
}

// buildKnownContracts lists the contracts whose admin functions the multisig
// is expected to call. ERC20 is deliberately left out.
func buildKnownContracts(address common.Address, peers Peers) []knownContract {
	specs := []struct {
		name     string
		address  common.Address
		metadata *bind.MetaData
	}{
		{"Payment Processor", peers.PaymentProcessor, &intermediatedpaymentprocessor.IntermediatedpaymentprocessorMetaData},
		{"Simple Payment Processor", peers.SimplePaymentProcessor, &simplepaymentprocessor.SimplepaymentprocessorMetaData},
		{"Payment Processor Storage", peers.PaymentProcessorStorage, &paymentprocessorstorage.PaymentprocessorstorageMetaData},
		{"Multisig", address, &gen.MultisigMetaData},
	}

	contracts := make([]knownContract, 0, len(specs))
	for _, spec := range specs {
		parsed, err := spec.metadata.ParseABI()
		if err != nil {
			slog.Error("abi parse failed for multisig decoding", "contract", spec.name, "error", err)
			continue
		}
		contracts = append(contracts, knownContract{
			name:    spec.name,
			address: spec.address,
			abi:     parsed,
		})
	}
	return contracts
}

type argValue struct {
	Name  string
	Value string
}

// action is a decoded multisig payload: which function is being called, in
// which contract, and with what arguments.
type action struct {
	Name     string
	Contract string
	Args     []argValue
}

// decodeAction resolves the calldata of a multisig transaction against the
// known contract ABIs. Contracts deployed at the target address are tried
// first so shared selectors resolve to the right contract.
func (c *Multisig) decodeAction(target common.Address, data []byte) *action {
	if len(data) < 4 {
		return nil
	}

	ordered := make([]knownContract, 0, len(c.known))
	for _, kc := range c.known {
		if kc.address == target {
			ordered = append(ordered, kc)
		}
	}
	for _, kc := range c.known {
		if kc.address != target {
			ordered = append(ordered, kc)
		}
	}

	for _, kc := range ordered {
		for _, method := range kc.abi.Methods {
			if !bytes.Equal(method.ID, data[:4]) {
				continue
			}

			decoded := &action{
				Name:     humanize(method.Name),
				Contract: kc.name,
			}

			values, err := method.Inputs.Unpack(data[4:])
			if err != nil {
				slog.Warn("calldata decode failed", "method", method.Name, "error", err)
				return decoded
			}
			for i, value := range values {
				name := method.Inputs[i].Name
				if name == "" {
					name = fmt.Sprintf("Argument %d", i+1)
				}
				decoded.Args = append(decoded.Args, argValue{
					Name:  humanize(name),
					Value: c.formatArg(value),
				})
			}
			return decoded
		}
	}
	return nil
}

// humanize turns identifiers like "setFeeRate" or "_newFeeRate" into
// "Set Fee Rate" / "New Fee Rate".
func humanize(identifier string) string {
	identifier = strings.TrimLeft(identifier, "_")
	var b strings.Builder
	for i, r := range identifier {
		if i == 0 {
			b.WriteRune(unicode.ToUpper(r))
			continue
		}
		if unicode.IsUpper(r) || unicode.IsDigit(r) {
			prev := rune(identifier[i-1])
			if !unicode.IsUpper(prev) && !unicode.IsDigit(prev) {
				b.WriteRune(' ')
			}
		}
		b.WriteRune(r)
	}
	return b.String()
}

func (c *Multisig) formatArg(value any) string {
	switch v := value.(type) {
	case common.Address:
		return c.addressLink(v)
	case *big.Int:
		return v.String()
	case [32]byte:
		return "`" + common.Hash(v).Hex() + "`"
	case []byte:
		return "`" + truncate(hexutil.Encode(v), 120) + "`"
	case bool:
		if v {
			return "yes"
		}
		return "no"
	case string:
		return v
	default:
		return truncate(fmt.Sprintf("%+v", v), 200)
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "…"
}
