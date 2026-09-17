package multisig

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var (
	storageAddress  = common.HexToAddress("0x2222222222222222222222222222222222222222")
	multisigAddress = common.HexToAddress("0x4444444444444444444444444444444444444444")
	someAccount     = common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
)

func adminMultisig() *Multisig {
	return &Multisig{
		explorerURL: "https://basescan.org",
		known: buildKnownContracts(multisigAddress, Peers{
			PaymentProcessor:        common.HexToAddress("0x1111111111111111111111111111111111111111"),
			SimplePaymentProcessor:  common.HexToAddress("0x3333333333333333333333333333333333333333"),
			PaymentProcessorStorage: storageAddress,
		}),
	}
}

func packFor(t *testing.T, c *Multisig, target common.Address, method string, args ...any) []byte {
	t.Helper()

	for _, kc := range c.known {
		if kc.address != target {
			continue
		}
		if _, ok := kc.abi.Methods[method]; !ok {
			t.Fatalf("%s is not in the %s ABI", method, kc.name)
		}
		packed, err := kc.abi.Pack(method, args...)
		if err != nil {
			t.Fatalf("packing %s: %v", method, err)
		}
		return packed
	}
	t.Fatalf("no known contract at %s", target)
	return nil
}

func TestDecodeAdminActions(t *testing.T) {
	var txID [32]byte
	txID[31] = 7

	tests := []struct {
		name         string
		target       common.Address
		method       string
		args         []any
		wantTitle    string
		wantContract string
		wantArgs     []string
	}{
		{
			name: "set fee signer", target: storageAddress, method: "setFeeSigner",
			args: []any{someAccount}, wantTitle: "Set Fee Signer",
			wantContract: "Payment Processor Storage", wantArgs: []string{"Fee Signer", "0x5aAe…eAed"},
		},
		{
			name: "set intermediated platforms operator", target: storageAddress,
			method: "setIntermediatedPlatformsOperator", args: []any{someAccount},
			wantTitle: "Set Intermediated Platforms Operator", wantContract: "Payment Processor Storage",
			wantArgs: []string{"0x5aAe…eAed"},
		},
		{
			name: "set emergency pauser", target: storageAddress, method: "setEmergencyPauser",
			args: []any{someAccount}, wantTitle: "Set Emergency Pauser",
			wantContract: "Payment Processor Storage", wantArgs: []string{"Emergency Pauser", "0x5aAe…eAed"},
		},
		{
			name: "pause", target: storageAddress, method: "pause",
			wantTitle: "Pause", wantContract: "Payment Processor Storage",
		},
		{
			name: "unpause", target: storageAddress, method: "unpause",
			wantTitle: "Unpause", wantContract: "Payment Processor Storage",
		},
		{
			name: "approve emergency pause", target: storageAddress, method: "approveEmergencyPause",
			wantTitle: "Approve Emergency Pause", wantContract: "Payment Processor Storage",
		},
		{
			name: "transfer ownership", target: storageAddress, method: "transferOwnership",
			args: []any{someAccount}, wantTitle: "Transfer Ownership",
			wantContract: "Payment Processor Storage", wantArgs: []string{"New Owner", "0x5aAe…eAed"},
		},
		{
			name: "cancel transaction", target: multisigAddress, method: "cancelTransaction",
			args: []any{txID}, wantTitle: "Cancel Transaction", wantContract: "Multisig",
			wantArgs: []string{common.Hash(txID).Hex()},
		},
		{
			name: "add signer", target: multisigAddress, method: "addSigner",
			args: []any{someAccount}, wantTitle: "Add Signer", wantContract: "Multisig",
			wantArgs: []string{"Signer", "0x5aAe…eAed"},
		},
		{
			name: "remove signer", target: multisigAddress, method: "removeSigner",
			args: []any{someAccount}, wantTitle: "Remove Signer", wantContract: "Multisig",
			wantArgs: []string{"Signer", "0x5aAe…eAed"},
		},
		{
			name: "update threshold", target: multisigAddress, method: "updateThreshold",
			args: []any{big.NewInt(3)}, wantTitle: "Update Threshold", wantContract: "Multisig",
			wantArgs: []string{"New Threshold", "**3**"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := adminMultisig()

			act := c.decodeAction(tt.target, packFor(t, c, tt.target, tt.method, tt.args...))

			if act == nil {
				t.Fatalf("%s did not decode", tt.method)
			}
			if actionTitle(act) != tt.wantTitle {
				t.Errorf("title = %q, want %q", actionTitle(act), tt.wantTitle)
			}
			if act.Contract != tt.wantContract {
				t.Errorf("contract = %q, want %q", act.Contract, tt.wantContract)
			}
			if len(act.Args) != len(tt.args) {
				t.Errorf("got %d args, want %d", len(act.Args), len(tt.args))
			}

			lines := strings.Join(actionArgLines(act), "\n")
			for _, want := range tt.wantArgs {
				if !strings.Contains(lines, want) {
					t.Errorf("arg lines %q, want them to contain %q", lines, want)
				}
			}
		})
	}
}

func TestAdminActionsNameTheirTargetContract(t *testing.T) {
	c := adminMultisig()

	storageCall := c.decodeAction(storageAddress, packFor(t, c, storageAddress, "pause"))
	if got := c.targetName(storageCall, storageAddress); !strings.Contains(got, "Payment Processor Storage") {
		t.Errorf("targetName = %q, want it to name the storage contract", got)
	}

	msigCall := c.decodeAction(multisigAddress, packFor(t, c, multisigAddress, "addSigner", someAccount))
	if got := c.targetName(msigCall, multisigAddress); !strings.Contains(got, "Multisig") {
		t.Errorf("targetName = %q, want it to name the multisig", got)
	}
}

func TestNoArgumentAdminActionsRenderNoArgumentBlock(t *testing.T) {
	c := adminMultisig()

	for _, method := range []string{"pause", "unpause", "approveEmergencyPause"} {
		act := c.decodeAction(storageAddress, packFor(t, c, storageAddress, method))

		if act == nil {
			t.Fatalf("%s did not decode", method)
		}
		if lines := actionArgLines(act); lines != nil {
			t.Errorf("%s produced argument lines %q, want none", method, lines)
		}
	}
}
