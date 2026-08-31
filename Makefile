INTERMEDIATED_ABI_FOLDER = internal/blockchain/gen/IntermediatedPaymentProcessor
INTERMEDIATED_ABI_FILE   = $(INTERMEDIATED_ABI_FOLDER)/IntermediatedPaymentProcessor.json
INTERMEDIATED_GO_OUT     = $(INTERMEDIATED_ABI_FOLDER)/IntermediatedPaymentProcessor.go
INTERMEDIATED_GO_PKG     = intermediatedprocessor

STORAGE_ABI_FOLDER  = internal/blockchain/gen/PaymentProcessorStorage
STORAGE_ABI_FILE    = $(STORAGE_ABI_FOLDER)/PaymentProcessorStorage.json
STORAGE_GO_OUT      = $(STORAGE_ABI_FOLDER)/PaymentProcessorStorage.go
STORAGE_GO_PKG      = processorstorage

SIMPLE_ABI_FOLDER  = internal/blockchain/gen/SimplePaymentProcessor
SIMPLE_ABI_FILE    = $(SIMPLE_ABI_FOLDER)/SimplePaymentProcessor.json
SIMPLE_GO_OUT      = $(SIMPLE_ABI_FOLDER)/SimplePaymentProcessor.go
SIMPLE_GO_PKG      = simpleprocessor

ERC20_ABI_FOLDER  = internal/blockchain/gen/ERC20
ERC20_ABI_FILE    = $(ERC20_ABI_FOLDER)/ERC20.json
ERC20_GO_OUT      = $(ERC20_ABI_FOLDER)/ERC20.go
ERC20_GO_PKG      = erc20

AUTOMATION_ABI_FOLDER  = internal/blockchain/gen/PaymentAutomation
AUTOMATION_ABI_FILE    = $(AUTOMATION_ABI_FOLDER)/PaymentAutomation.json
AUTOMATION_GO_OUT      = $(AUTOMATION_ABI_FOLDER)/PaymentAutomation.go
AUTOMATION_GO_PKG      = paymentautomation

NOTES_ABI_FOLDER  = internal/blockchain/gen/Notes
NOTES_ABI_FILE    = $(NOTES_ABI_FOLDER)/Notes.json
NOTES_GO_OUT      = $(NOTES_ABI_FOLDER)/Notes.go
NOTES_GO_PKG      = notescontract

.PHONY: clean gen run

clean:
	@rm -f $(INTERMEDIATED_GO_OUT) $(STORAGE_GO_OUT) $(SIMPLE_GO_OUT) $(ERC20_GO_OUT) $(AUTOMATION_GO_OUT) $(NOTES_GO_OUT)

gen:
	@abigen --v2 --abi $(INTERMEDIATED_ABI_FILE) --pkg $(INTERMEDIATED_GO_PKG) --out $(INTERMEDIATED_GO_OUT)
	@abigen --v2 --abi $(SIMPLE_ABI_FILE) --pkg $(SIMPLE_GO_PKG) --out $(SIMPLE_GO_OUT)
	@abigen --v2 --abi $(STORAGE_ABI_FILE) --pkg $(STORAGE_GO_PKG) --out $(STORAGE_GO_OUT)
	@abigen --v2 --abi $(ERC20_ABI_FILE) --pkg $(ERC20_GO_PKG) --out $(ERC20_GO_OUT)
	@abigen --v2 --abi $(AUTOMATION_ABI_FILE) --pkg $(AUTOMATION_GO_PKG) --out $(AUTOMATION_GO_OUT)
	@abigen --v2 --abi $(NOTES_ABI_FILE) --pkg $(NOTES_GO_PKG) --out $(NOTES_GO_OUT)

run:
	@go run ./cmd/server