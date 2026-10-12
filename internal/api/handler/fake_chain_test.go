package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/simplepaymentprocessor"
	ippgen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/intermediatedpaymentprocessor"
	sppgen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/simplepaymentprocessor"
)

var (
	fakeIntermediatedAddress = common.HexToAddress("0x1000000000000000000000000000000000000001")
	fakeSimpleAddress        = common.HexToAddress("0x2000000000000000000000000000000000000002")
)

type fakeChain struct {
	intermediated map[int64]uint8
	meta          map[int64][]int64
	simple        map[int64]uint8
	failWith      string
}

func readyChain() *fakeChain {
	return &fakeChain{
		intermediated: map[int64]uint8{
			42: intermediatedpaymentprocessor.StateCreated,
			43: intermediatedpaymentprocessor.StatePaid,
			44: intermediatedpaymentprocessor.StateCreated,
		},
		meta:   map[int64][]int64{42: {43, 44}},
		simple: map[int64]uint8{42: simplepaymentprocessor.StatePaid},
	}
}

func (f *fakeChain) install(t *testing.T, h *ContractHandler) {
	t.Helper()

	intermediatedABI, err := ippgen.IntermediatedpaymentprocessorMetaData.ParseABI()
	if err != nil {
		t.Fatalf("parsing the intermediated ABI: %v", err)
	}
	simpleABI, err := sppgen.SimplepaymentprocessorMetaData.ParseABI()
	if err != nil {
		t.Fatalf("parsing the simple ABI: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage   `json:"id"`
			Method string            `json:"method"`
			Params []json.RawMessage `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reply := map[string]any{"jsonrpc": "2.0", "id": req.ID}
		if result, err := f.call(intermediatedABI, simpleABI, req.Method, req.Params); err != nil {
			reply["error"] = map[string]any{"code": -32000, "message": err.Error()}
		} else {
			reply["result"] = result
		}
		json.NewEncoder(w).Encode(reply)
	}))
	t.Cleanup(server.Close)

	rpc, err := ethclient.Dial(server.URL)
	if err != nil {
		t.Fatalf("dialing the fake chain: %v", err)
	}
	t.Cleanup(rpc.Close)

	client := &blockchain.Client{HTTP: rpc}
	h.PaymentProcessor = intermediatedpaymentprocessor.NewPaymentprocessor(client, fakeIntermediatedAddress, "", nil)
	h.SimplePaymentProcessor = simplepaymentprocessor.NewSimplePaymentProcessor(client, fakeSimpleAddress)
}

func (f *fakeChain) call(intermediatedABI, simpleABI *abi.ABI, method string, params []json.RawMessage) (string, error) {
	if f.failWith != "" {
		return "", errors.New(f.failWith)
	}
	if method != "eth_call" || len(params) == 0 {
		return "", fmt.Errorf("unexpected call %s", method)
	}

	var msg struct {
		To    common.Address `json:"to"`
		Data  hexutil.Bytes  `json:"data"`
		Input hexutil.Bytes  `json:"input"`
	}
	if err := json.Unmarshal(params[0], &msg); err != nil {
		return "", err
	}
	data := msg.Input
	if len(data) == 0 {
		data = msg.Data
	}
	if len(data) < 4 {
		return "", errors.New("call data too short")
	}

	contractABI := intermediatedABI
	if msg.To == fakeSimpleAddress {
		contractABI = simpleABI
	}
	m, err := contractABI.MethodById(data[:4])
	if err != nil {
		return "", err
	}
	args, err := m.Inputs.Unpack(data[4:])
	if err != nil {
		return "", err
	}
	id := args[0].(*big.Int).Int64()

	var out []byte
	switch m.Name {
	case "getInvoice":
		out, err = m.Outputs.Pack(withZeroBigInts(ippgen.IIntermediatedPaymentProcessorInvoice{State: f.intermediated[id]}))
	case "getMetaInvoice":
		subInvoiceIDs := []*big.Int{}
		for _, sub := range f.meta[id] {
			subInvoiceIDs = append(subInvoiceIDs, big.NewInt(sub))
		}
		out, err = m.Outputs.Pack(ippgen.IIntermediatedPaymentProcessorMetaInvoice{Price: big.NewInt(0), SubInvoiceIds: subInvoiceIDs})
	case "getInvoiceData":
		out, err = m.Outputs.Pack(withZeroBigInts(sppgen.ISimplePaymentProcessorInvoice{State: f.simple[id]}))
	default:
		return "", fmt.Errorf("unexpected method %s", m.Name)
	}
	if err != nil {
		return "", err
	}
	return hexutil.Encode(out), nil
}

func withZeroBigInts[T any](v T) T {
	rv := reflect.ValueOf(&v).Elem()
	bigInt := reflect.TypeOf((*big.Int)(nil))
	for i := 0; i < rv.NumField(); i++ {
		if field := rv.Field(i); field.Type() == bigInt && field.IsNil() {
			field.Set(reflect.ValueOf(big.NewInt(0)))
		}
	}
	return v
}
