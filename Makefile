INTERMEDIATED_ABI_FOLDER = internal/blockchain/gen/intermediatedpaymentprocessor
INTERMEDIATED_ABI_FILE   = $(INTERMEDIATED_ABI_FOLDER)/IntermediatedPaymentProcessor.json
INTERMEDIATED_GO_OUT     = $(INTERMEDIATED_ABI_FOLDER)/intermediated_payment_processor.go
INTERMEDIATED_GO_PKG     = intermediatedpaymentprocessor

STORAGE_ABI_FOLDER  = internal/blockchain/gen/paymentprocessorstorage
STORAGE_ABI_FILE    = $(STORAGE_ABI_FOLDER)/PaymentProcessorStorage.json
STORAGE_GO_OUT      = $(STORAGE_ABI_FOLDER)/payment_processor_storage.go
STORAGE_GO_PKG      = paymentprocessorstorage

SIMPLE_ABI_FOLDER  = internal/blockchain/gen/simplepaymentprocessor
SIMPLE_ABI_FILE    = $(SIMPLE_ABI_FOLDER)/SimplePaymentProcessor.json
SIMPLE_GO_OUT      = $(SIMPLE_ABI_FOLDER)/simple_payment_processor.go
SIMPLE_GO_PKG      = simplepaymentprocessor

ERC20_ABI_FOLDER  = internal/blockchain/gen/erc20
ERC20_ABI_FILE    = $(ERC20_ABI_FOLDER)/ERC20.json
ERC20_GO_OUT      = $(ERC20_ABI_FOLDER)/erc20.go
ERC20_GO_PKG      = erc20

AUTOMATION_ABI_FOLDER  = internal/blockchain/gen/paymentautomation
AUTOMATION_ABI_FILE    = $(AUTOMATION_ABI_FOLDER)/PaymentAutomation.json
AUTOMATION_GO_OUT      = $(AUTOMATION_ABI_FOLDER)/payment_automation.go
AUTOMATION_GO_PKG      = paymentautomation

ORACLE_ABI_FOLDER  = internal/blockchain/gen/oraclemanager
ORACLE_ABI_FILE    = $(ORACLE_ABI_FOLDER)/OracleManager.json
ORACLE_GO_OUT      = $(ORACLE_ABI_FOLDER)/oracle_manager.go
ORACLE_GO_PKG      = oraclemanager

NOTES_ABI_FOLDER  = internal/blockchain/gen/notes
NOTES_ABI_FILE    = $(NOTES_ABI_FOLDER)/Notes.json
NOTES_GO_OUT      = $(NOTES_ABI_FOLDER)/notes.go
NOTES_GO_PKG      = notes

.PHONY: clean gen run

clean:
	@rm -f $(ORACLE_GO_OUT) $(INTERMEDIATED_GO_OUT) $(STORAGE_GO_OUT) $(SIMPLE_GO_OUT) $(ERC20_GO_OUT) $(AUTOMATION_GO_OUT) $(NOTES_GO_OUT)

gen:
	@abigen --v2 --abi $(INTERMEDIATED_ABI_FILE) --pkg $(INTERMEDIATED_GO_PKG) --out $(INTERMEDIATED_GO_OUT)
	@abigen --v2 --abi $(SIMPLE_ABI_FILE) --pkg $(SIMPLE_GO_PKG) --out $(SIMPLE_GO_OUT)
	@abigen --v2 --abi $(STORAGE_ABI_FILE) --pkg $(STORAGE_GO_PKG) --out $(STORAGE_GO_OUT)
	@abigen --v2 --abi $(ERC20_ABI_FILE) --pkg $(ERC20_GO_PKG) --out $(ERC20_GO_OUT)
	@abigen --v2 --abi $(AUTOMATION_ABI_FILE) --pkg $(AUTOMATION_GO_PKG) --out $(AUTOMATION_GO_OUT)
	@abigen --v2 --abi $(NOTES_ABI_FILE) --pkg $(NOTES_GO_PKG) --out $(NOTES_GO_OUT)
	@abigen --v2 --abi $(ORACLE_ABI_FILE) --pkg $(ORACLE_GO_PKG) --out $(ORACLE_GO_OUT)

run:
	@go run ./cmd/server