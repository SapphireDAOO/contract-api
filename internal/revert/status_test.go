package revert

import (
	"net/http"
	"testing"
)

func TestStatusCodesKeysAreDescriptions(t *testing.T) {
	descriptions := make(map[string]bool, len(Descriptions))
	for _, description := range Descriptions {
		descriptions[description] = true
	}

	for reason := range StatusCodes {
		if !descriptions[reason] {
			t.Errorf("StatusCodes has %q, which no selector in Descriptions produces", reason)
		}
	}
}

func TestStatusCodesAreClientErrorsOrUnavailable(t *testing.T) {
	for reason, code := range StatusCodes {
		if code >= 400 && code < 500 {
			continue
		}
		if code == http.StatusServiceUnavailable {
			continue
		}
		t.Errorf("StatusCodes[%q] = %d, want a 4xx status or 503", reason, code)
	}
}

func TestStatusCodeLookupFallsBackToServerError(t *testing.T) {
	if code := StatusCodes[Reason(revertWith("0xdeadbeef"))]; code != 0 {
		t.Errorf("unmapped reason has status %d, want 0 so callers can default", code)
	}
	if code := StatusCodes["The specified invoice does not exist."]; code != http.StatusNotFound {
		t.Errorf("missing invoice status = %d, want %d", code, http.StatusNotFound)
	}
}
