package crypto

import (
	"encoding/base64"
	"errors"

	"github.com/gtank/cryptopasta"
)

// EncryptStringToBase64 encrypts the plaintext with the provided key and returns the base64-encoded result.
func EncryptStringToBase64(plaintext string, key *AES256Key) (base64Ciphertext string, err error) {

	if key == nil {
		err = errors.New("tried to encrypt with nil key")
		return
	}

	ciphertext, err := cryptopasta.Encrypt([]byte(plaintext), (*[AES256KeyLengthInBytes]byte)(key))
	if err != nil {
		return
	}

	base64Ciphertext = base64.StdEncoding.EncodeToString(ciphertext)
	return
}

// DecryptStringFromBase64 decodes the provided base64 string (e.g. from a previous call to EncryptStringToBase64),
// decrypts the result with the provided key, and returns the resulting string.
func DecryptStringFromBase64(base64Ciphertext string, key *AES256Key) (plaintext string, err error) {

	if key == nil {
		err = errors.New("tried to decrypt with nil key")
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		return
	}

	plainBytes, err := cryptopasta.Decrypt(ciphertext, (*[AES256KeyLengthInBytes]byte)(key))
	if err != nil {
		return
	}

	plaintext = string(plainBytes)
	return
}
