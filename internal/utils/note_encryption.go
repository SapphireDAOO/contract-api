package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const noteIVLength = 16

func notesSecretKey() string {
	return os.Getenv("NOTES_SECRET_KEY")
}

func noteCipherKey(secretKey string) []byte {
	sum := sha256.Sum256([]byte(secretKey))
	return sum[:]
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	padding := int(data[len(data)-1])
	if padding == 0 || padding > blockSize || padding > len(data) {
		return nil, errors.New("invalid padding")
	}

	for _, b := range data[len(data)-padding:] {
		if int(b) != padding {
			return nil, errors.New("invalid padding")
		}
	}

	return data[:len(data)-padding], nil
}

// EncryptNote produces the "<base64 iv>:<base64 ciphertext>" payload.
func EncryptNote(note, secretKey string) (string, error) {
	block, err := aes.NewCipher(noteCipherKey(secretKey))
	if err != nil {
		return "", err
	}

	iv := make([]byte, noteIVLength)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate iv: %w", err)
	}

	padded := pkcs7Pad([]byte(note), block.BlockSize())
	ciphertext := make([]byte, len(padded))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, padded)

	return base64.StdEncoding.EncodeToString(iv) + ":" +
		base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptNote(encryptedNote, secretKey string) (string, error) {
	ivBase64, encrypted, found := strings.Cut(encryptedNote, ":")
	if !found || ivBase64 == "" || encrypted == "" {
		return encryptedNote, nil
	}

	iv, err := base64.StdEncoding.DecodeString(ivBase64)
	if err != nil {
		return "", fmt.Errorf("invalid iv: %w", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("invalid ciphertext: %w", err)
	}

	block, err := aes.NewCipher(noteCipherKey(secretKey))
	if err != nil {
		return "", err
	}

	if len(iv) != block.BlockSize() {
		return "", errors.New("invalid iv length")
	}

	if len(ciphertext) == 0 || len(ciphertext)%block.BlockSize() != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(plaintext, ciphertext)

	unpadded, err := pkcs7Unpad(plaintext, block.BlockSize())
	if err != nil {
		return "", err
	}

	return string(unpadded), nil
}

func ToEncryptedNoteBytes(note string) ([]byte, error) {
	trimmed := strings.TrimSpace(note)
	if trimmed == "" {
		return []byte{}, nil
	}

	secretKey := notesSecretKey()
	if secretKey == "" {
		return []byte(trimmed), nil
	}

	payload, err := EncryptNote(trimmed, secretKey)
	if err != nil {
		return nil, err
	}

	return []byte(payload), nil
}

func DecryptNoteBlob(blob []byte) (string, error) {
	encrypted := strings.TrimSpace(string(blob))
	if encrypted == "" {
		return "", nil
	}

	secretKey := notesSecretKey()
	if secretKey == "" || !strings.Contains(encrypted, ":") {
		return encrypted, nil
	}

	return DecryptNote(encrypted, secretKey)
}
