package revert

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func abiErrorSelectors(t *testing.T) map[string]string {
	t.Helper()

	paths, err := filepath.Glob(filepath.Join("..", "blockchain", "gen", "*", "*.json"))
	if err != nil {
		t.Fatalf("globbing the ABIs: %v", err)
	}
	if len(paths) == 0 {
		t.Skip("no generated ABIs found")
	}

	selectors := make(map[string]string)
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			t.Fatalf("opening %s: %v", path, err)
		}
		parsed, err := abi.JSON(f)
		f.Close()
		if err != nil {
			t.Fatalf("parsing %s: %v", path, err)
		}
		for _, e := range parsed.Errors {
			selectors["0x"+hex.EncodeToString(e.ID.Bytes()[:4])] = e.Sig
		}
	}
	return selectors
}

func TestEveryABIErrorHasADescription(t *testing.T) {
	for selector, sig := range abiErrorSelectors(t) {
		if _, ok := Descriptions[selector]; !ok {
			t.Errorf("no description for %s (%s); the ABIs and Descriptions have drifted", selector, sig)
		}
	}
}

func TestDescriptionsHaveNoStaleSelectors(t *testing.T) {
	selectors := abiErrorSelectors(t)

	for selector := range Descriptions {
		if _, ok := selectors[selector]; !ok {
			t.Errorf("%s is described but is in no ABI", selector)
		}
	}
}

func TestEveryDescriptionIsNonEmpty(t *testing.T) {
	for selector, description := range Descriptions {
		if strings.TrimSpace(description) == "" {
			t.Errorf("selector %s has an empty description", selector)
		}
	}
}

func TestDescriptionKeysAreFourByteSelectors(t *testing.T) {
	for selector, description := range Descriptions {
		if len(selector) != 10 || selector[:2] != "0x" {
			t.Errorf("selector %q is not a 0x-prefixed four-byte selector", selector)
		}
		if description == "" {
			t.Errorf("selector %q has an empty description", selector)
		}
	}
}

func TestUnsupportedTokenIsDescribed(t *testing.T) {
	if _, ok := Descriptions[UnsupportedToken]; !ok {
		t.Errorf("Descriptions is missing the UnsupportedToken selector %q", UnsupportedToken)
	}
}
