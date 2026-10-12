package feereceiver

import (
	"fmt"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/config"
	pb "github.com/SapphireDAOO/contract-api/proto/feereceiverpb"
)

func ParseProcessorKind(value string) (pb.ProcessorKind, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "simple":
		return pb.ProcessorKind_PROCESSOR_KIND_SIMPLE, nil
	case "intermediated":
		return pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED, nil
	default:
		return pb.ProcessorKind_PROCESSOR_KIND_UNSPECIFIED,
			fmt.Errorf("%w, got %q", ErrInvalidProcessor, value)
	}
}

func ParseInvoiceKind(value string) (pb.InvoiceKind, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "single":
		return pb.InvoiceKind_INVOICE_KIND_SINGLE, nil
	case "meta":
		return pb.InvoiceKind_INVOICE_KIND_META, nil
	default:
		return pb.InvoiceKind_INVOICE_KIND_UNSPECIFIED,
			fmt.Errorf("%w, got %q", ErrInvalidKind, value)
	}
}

func ResolveFeeToken(tokens config.Tokens, symbol string) (string, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return "", nil
	}

	address, ok := tokens.Address(symbol)
	if !ok {
		return "", fmt.Errorf("%w %s (known: %s)",
			ErrUnknownToken, symbol, strings.Join(tokens.Symbols(), ", "))
	}
	return address.Hex(), nil
}
