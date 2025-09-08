ADVANCED_ABI_FOLDER = internal/blockchain/gen/AdvancedPaymentProcessor
ADVANCED_ABI_FILE   = $(ADVANCED_ABI_FOLDER)/AdvancedPaymentProcessor.json
ADVANCED_GO_OUT     = $(ADVANCED_ABI_FOLDER)/AdvancedPaymentProcessor.go
ADVANCED_GO_PKG     = advancedprocessor

STORAGE_ABI_FOLDER  = internal/blockchain/gen/PaymentProcessorStorage
STORAGE_ABI_FILE    = $(STORAGE_ABI_FOLDER)/PaymentProcessorStorage.json
STORAGE_GO_OUT      = $(STORAGE_ABI_FOLDER)/PaymentProcessorStorage.go
STORAGE_GO_PKG      = processorstorage

clean:
	@rm -f $(ADVANCED_GO_OUT) $(STORAGE_GO_OUT)

gen:
	@abigen --v2 --abi $(ADVANCED_ABI_FILE) --pkg $(ADVANCED_GO_PKG) --out $(ADVANCED_GO_OUT)
	@abigen --v2 --abi $(STORAGE_ABI_FILE) --pkg $(STORAGE_GO_PKG) --out $(STORAGE_GO_OUT)

run:
	@go run ./server/main.go