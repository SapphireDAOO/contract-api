package api

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const TX_URL string = "https://sepolia.etherscan.io/tx/"

type ContractHandler struct {
	Contract *blockchain.Contract
	baseUrl  string
}

func NewContractHandler(contract *blockchain.Contract, baseUrl string) *ContractHandler {
	return &ContractHandler{Contract: contract, baseUrl: baseUrl}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam

	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if len(invoice) == 0 {
		utils.Error(w, http.StatusBadRequest, nil, "no invoice parameters provided")
		return
	}

	var token string
	var response any
	var err error

	if len(invoice) == 1 {
		var res *blockchain.SingleInvoiceResponse
		res, err = h.Contract.CreateInvoice(invoice)
		if err == nil {
			token, err = utils.GenerateToken(*res.Key)
			response = res
		}
		res.Url = h.baseUrl + token
	} else {
		var res *blockchain.MetaInvoiceResponse
		res, err = h.Contract.CreateInvoices(invoice)
		if err == nil {
			token, err = utils.GenerateToken(*res.Key)
			response = res
		}
		res.Url = h.baseUrl + token
	}

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "error processing invoice")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ContractHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Id common.Hash `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.Contract.Cancel(input.Id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"hash":   TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) Refund(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Id common.Hash `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.Contract.Refund(input.Id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"hash":   TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) Release(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Id common.Hash `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.Contract.Release(input.Id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"hash":   TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) HandleDispute(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId      string                       `json:"orderId"`
		Resolution   blockchain.MarketplaceAction `json:"resolution"`
		Resolver     *common.Address              `json:"resolver"`
		SellersShare *big.Int                     `json:"sellersShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	var id common.Hash

	if len(input.OrderId) == 66 && strings.HasPrefix(input.OrderId, "0x") {
		id = common.HexToHash(input.OrderId)
	} else {
		hashed, err := utils.Keccak256(input.OrderId)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err, "failed to generate order ID hash")
			return
		}
		id = *hashed
	}

	txHash, err := h.Contract.ReleaseEscrow(id, input.Resolver, input.Resolution, input.SellersShare)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"hash":   TX_URL + txHash.Hex(),
	})
}
