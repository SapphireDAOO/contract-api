ABI_FOLDER = PaymentProcessor
ABI_FILE = abi/abi.json
GO_OUT = $(ABI_FOLDER)/AdvancedPaymentProcessor.go
GO_PKG = paymentprocessor

gen:
	@mkdir -p $(ABI_FOLDER)
	@abigen --abi $(ABI_FILE) --pkg $(GO_PKG) --out $(GO_OUT)