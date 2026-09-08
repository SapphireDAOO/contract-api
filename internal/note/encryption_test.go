package note

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"strings"
	"testing"
)

const testSecret = "test-secret-key"

func TestEncryptDecryptRoundTrip(t *testing.T) {
	notes := []string{
		"a",
		"hello",
		"exactly 16 bytes",
		"seventeen bytes!!",
		strings.Repeat("x", 1024),
		"unicode ☃ note",
		"has:a:colon",
		"trailing spaces   ",
		"\x00\x01\x02 control",
	}

	for _, want := range notes {
		t.Run(want, func(t *testing.T) {
			encrypted, err := EncryptNote(want, testSecret)
			if err != nil {
				t.Fatalf("EncryptNote returned %v", err)
			}

			got, err := DecryptNote(encrypted, testSecret)
			if err != nil {
				t.Fatalf("DecryptNote returned %v", err)
			}
			if got != want {
				t.Errorf("round trip produced %q, want %q", got, want)
			}
		})
	}
}

func TestEncryptDecryptEmptyNote(t *testing.T) {
	encrypted, err := EncryptNote("", testSecret)
	if err != nil {
		t.Fatalf("EncryptNote returned %v", err)
	}

	got, err := DecryptNote(encrypted, testSecret)
	if err != nil {
		t.Fatalf("DecryptNote returned %v", err)
	}
	if got != "" {
		t.Errorf("round trip of the empty note produced %q", got)
	}
}

func TestEncryptNotePayloadShape(t *testing.T) {
	encrypted, err := EncryptNote("hello", testSecret)
	if err != nil {
		t.Fatalf("EncryptNote returned %v", err)
	}

	ivPart, cipherPart, found := strings.Cut(encrypted, ":")
	if !found {
		t.Fatalf("payload %q is not <iv>:<ciphertext>", encrypted)
	}

	iv, err := base64.StdEncoding.DecodeString(ivPart)
	if err != nil {
		t.Fatalf("iv is not base64: %v", err)
	}
	if len(iv) != noteIVLength {
		t.Errorf("iv is %d bytes, want %d", len(iv), noteIVLength)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cipherPart)
	if err != nil {
		t.Fatalf("ciphertext is not base64: %v", err)
	}
	if len(ciphertext) == 0 || len(ciphertext)%aes.BlockSize != 0 {
		t.Errorf("ciphertext is %d bytes, want a non-zero multiple of %d", len(ciphertext), aes.BlockSize)
	}
}

func TestEncryptNoteUsesAFreshIV(t *testing.T) {
	const note = "same note"

	first, err := EncryptNote(note, testSecret)
	if err != nil {
		t.Fatalf("EncryptNote returned %v", err)
	}
	second, err := EncryptNote(note, testSecret)
	if err != nil {
		t.Fatalf("EncryptNote returned %v", err)
	}

	if first == second {
		t.Error("encrypting the same note twice produced identical payloads")
	}

	for _, payload := range []string{first, second} {
		got, err := DecryptNote(payload, testSecret)
		if err != nil {
			t.Fatalf("DecryptNote returned %v", err)
		}
		if got != note {
			t.Errorf("DecryptNote = %q, want %q", got, note)
		}
	}
}

func TestDecryptNoteWithTheWrongKey(t *testing.T) {
	encrypted, err := EncryptNote("hello", testSecret)
	if err != nil {
		t.Fatalf("EncryptNote returned %v", err)
	}

	got, err := DecryptNote(encrypted, "a different secret")

	if err == nil && got == "hello" {
		t.Error("DecryptNote recovered the note with the wrong key")
	}
}

func TestDecryptNoteErrors(t *testing.T) {
	validIV := base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{1}, noteIVLength))
	validBlock := base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{2}, aes.BlockSize))

	tests := []struct {
		name    string
		payload string
		wantErr string
	}{
		{"invalid iv base64", "not base64!:" + validBlock, "invalid iv"},
		{"invalid ciphertext base64", validIV + ":not base64!", "invalid ciphertext"},
		{
			"short iv",
			base64.StdEncoding.EncodeToString([]byte{1, 2, 3}) + ":" + validBlock,
			"invalid iv length",
		},
		{
			"ciphertext is not a whole number of blocks",
			validIV + ":" + base64.StdEncoding.EncodeToString([]byte{1, 2, 3}),
			"not a multiple of the block size",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DecryptNote(tt.payload, testSecret)

			if err == nil {
				t.Fatalf("DecryptNote(%q) returned no error", tt.payload)
			}
			if tt.wantErr != "" && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestDecryptNotePassesThroughPlaintext(t *testing.T) {
	tests := []string{
		"a plain note",
		"",
		":no iv",

		"AQEBAQEBAQEBAQEBAQEBAQ==:",
	}

	for _, payload := range tests {
		t.Run(payload, func(t *testing.T) {
			got, err := DecryptNote(payload, testSecret)

			if err != nil {
				t.Fatalf("DecryptNote(%q) returned %v", payload, err)
			}
			if got != payload {
				t.Errorf("DecryptNote(%q) = %q, want it unchanged", payload, got)
			}
		})
	}
}

func TestPKCS7RoundTrip(t *testing.T) {
	for length := range 40 {
		data := bytes.Repeat([]byte{0xAB}, length)

		padded := pkcs7Pad(bytes.Clone(data), aes.BlockSize)
		if len(padded)%aes.BlockSize != 0 {
			t.Fatalf("padding %d bytes produced %d, not a multiple of %d", length, len(padded), aes.BlockSize)
		}

		if len(padded) <= length {
			t.Fatalf("padding %d bytes produced %d, want it to grow", length, len(padded))
		}

		unpadded, err := pkcs7Unpad(padded, aes.BlockSize)
		if err != nil {
			t.Fatalf("pkcs7Unpad after padding %d bytes returned %v", length, err)
		}
		if !bytes.Equal(unpadded, data) {
			t.Errorf("round trip of %d bytes produced %d bytes", length, len(unpadded))
		}
	}
}

func TestPKCS7UnpadRejectsBadInput(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{"empty", nil},
		{"not a whole number of blocks", bytes.Repeat([]byte{1}, aes.BlockSize+1)},
		{"zero padding byte", append(bytes.Repeat([]byte{1}, aes.BlockSize-1), 0)},
		{"padding longer than a block", append(bytes.Repeat([]byte{1}, aes.BlockSize-1), aes.BlockSize+1)},
		{"inconsistent padding bytes", append(bytes.Repeat([]byte{1}, aes.BlockSize-2), 2, 3)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := pkcs7Unpad(tt.data, aes.BlockSize); err == nil {
				t.Errorf("pkcs7Unpad(%v) returned no error", tt.data)
			}
		})
	}
}

func TestNoteCipherKeyIsAES256Sized(t *testing.T) {
	for _, secret := range []string{"", "short", strings.Repeat("long", 100)} {
		key := noteCipherKey(secret)

		if len(key) != 32 {
			t.Errorf("noteCipherKey(%q) is %d bytes, want 32", secret, len(key))
		}
		if _, err := aes.NewCipher(key); err != nil {
			t.Errorf("noteCipherKey(%q) is not a valid AES key: %v", secret, err)
		}
	}
}

func TestNoteCipherKeyIsDeterministicAndSecretDependent(t *testing.T) {
	if !bytes.Equal(noteCipherKey("a"), noteCipherKey("a")) {
		t.Error("noteCipherKey is not deterministic")
	}
	if bytes.Equal(noteCipherKey("a"), noteCipherKey("b")) {
		t.Error("noteCipherKey ignores the secret")
	}
}

func TestToEncryptedNoteBytes(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", testSecret)

	got, err := ToEncryptedNoteBytes("  a note  ")
	if err != nil {
		t.Fatalf("ToEncryptedNoteBytes returned %v", err)
	}

	if !strings.Contains(string(got), ":") {
		t.Fatalf("ToEncryptedNoteBytes produced %q, want an encrypted payload", got)
	}

	decrypted, err := DecryptNote(string(got), testSecret)
	if err != nil {
		t.Fatalf("DecryptNote returned %v", err)
	}
	if decrypted != "a note" {
		t.Errorf("decrypted %q, want the trimmed content", decrypted)
	}
}

func TestToEncryptedNoteBytesEmpty(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", testSecret)

	for _, content := range []string{"", "   ", "\t\n"} {
		got, err := ToEncryptedNoteBytes(content)
		if err != nil {
			t.Fatalf("ToEncryptedNoteBytes(%q) returned %v", content, err)
		}
		if len(got) != 0 {
			t.Errorf("ToEncryptedNoteBytes(%q) = %q, want no bytes", content, got)
		}
	}
}

func TestToEncryptedNoteBytesWithoutASecret(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", "")

	got, err := ToEncryptedNoteBytes("  a note  ")
	if err != nil {
		t.Fatalf("ToEncryptedNoteBytes returned %v", err)
	}
	if string(got) != "a note" {
		t.Errorf("ToEncryptedNoteBytes = %q, want the trimmed plaintext", got)
	}
}

func TestDecryptNoteBlob(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", testSecret)

	encrypted, err := ToEncryptedNoteBytes("a note")
	if err != nil {
		t.Fatalf("ToEncryptedNoteBytes returned %v", err)
	}

	got, err := DecryptNoteBlob(encrypted)
	if err != nil {
		t.Fatalf("DecryptNoteBlob returned %v", err)
	}
	if got != "a note" {
		t.Errorf("DecryptNoteBlob = %q, want %q", got, "a note")
	}
}

func TestDecryptNoteBlobEmpty(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", testSecret)

	for _, blob := range [][]byte{nil, {}, []byte("   ")} {
		got, err := DecryptNoteBlob(blob)
		if err != nil {
			t.Fatalf("DecryptNoteBlob(%q) returned %v", blob, err)
		}
		if got != "" {
			t.Errorf("DecryptNoteBlob(%q) = %q, want an empty string", blob, got)
		}
	}
}

func TestDecryptNoteBlobPassesThroughPlaintext(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", testSecret)

	got, err := DecryptNoteBlob([]byte("  a plain note  "))
	if err != nil {
		t.Fatalf("DecryptNoteBlob returned %v", err)
	}
	if got != "a plain note" {
		t.Errorf("DecryptNoteBlob = %q, want the trimmed plaintext", got)
	}
}

func TestDecryptNoteBlobWithoutASecret(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", "")

	payload := []byte("aXYaXYaXYaXY:Y2lwaGVy")

	got, err := DecryptNoteBlob(payload)
	if err != nil {
		t.Fatalf("DecryptNoteBlob returned %v", err)
	}
	if got != string(payload) {
		t.Errorf("DecryptNoteBlob = %q, want it unchanged", got)
	}
}

func TestNoteBlobRoundTrip(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", testSecret)

	notes := []string{"a", "hello there", strings.Repeat("x", 500), "unicode ☃"}

	for _, want := range notes {
		blob, err := ToEncryptedNoteBytes(want)
		if err != nil {
			t.Fatalf("ToEncryptedNoteBytes(%q) returned %v", want, err)
		}

		got, err := DecryptNoteBlob(blob)
		if err != nil {
			t.Fatalf("DecryptNoteBlob returned %v", err)
		}
		if got != want {
			t.Errorf("round trip of %q produced %q", want, got)
		}
	}
}
