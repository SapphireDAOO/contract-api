ABI_FOLDER = internal/blockchain/contracts/AdvancedPaymentProcessor
ABI_FILE = $(ABI_FOLDER)/AdvancedPaymentProcessor.json
GO_OUT = $(ABI_FOLDER)/AdvancedPaymentProcessor.go
GO_PKG = paymentprocessor


gen:
	@mkdir -p $(ABI_FOLDER)
	@abigen --v2  --abi $(ABI_FILE) --pkg $(GO_PKG) --out $(GO_OUT)

run:
	@go run ./server/main.go