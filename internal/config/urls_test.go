package config

import "testing"

func TestLink(t *testing.T) {
	tests := []struct {
		name        string
		explorerURL string
		path        string
		value       string
		want        string
	}{
		{
			name:        "transaction link",
			explorerURL: "https://etherscan.io",
			path:        "/tx/",
			value:       "0xabc",
			want:        "https://etherscan.io/tx/0xabc",
		},
		{
			name:        "address link",
			explorerURL: "https://etherscan.io",
			path:        "/address/",
			value:       "0xdef",
			want:        "https://etherscan.io/address/0xdef",
		},
		{
			name:        "no explorer falls back to the bare value",
			explorerURL: "",
			path:        "/tx/",
			value:       "0xabc",
			want:        "0xabc",
		},
		{
			name:        "no explorer ignores the path entirely",
			explorerURL: "",
			path:        "/address/",
			value:       "0xdef",
			want:        "0xdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Link(tt.explorerURL, tt.path, tt.value); got != tt.want {
				t.Errorf("Link(%q, %q, %q) = %q, want %q",
					tt.explorerURL, tt.path, tt.value, got, tt.want)
			}
		})
	}
}

func TestLinkWithAValidatedExplorer(t *testing.T) {
	urls := validURLs()

	if err := urls.validate(); err != nil {
		t.Fatalf("the fixture does not validate: %v", err)
	}

	got := Link(urls.Explorer, "/tx/", "0xabc")

	if got != "https://etherscan.io/tx/0xabc" {
		t.Errorf("Link = %q, want a single slash between host and path", got)
	}
}
