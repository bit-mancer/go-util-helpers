package crypto

import (
	"errors"

	"github.com/gtank/cryptopasta"
)

// Encrypt encrypts the plaintext with the provided key and returns the result.
func Encrypt(plaintext []byte, key *AES256Key) ([]byte, error) {

	if key == nil {
		return nil, errors.New("tried to encrypt with nil key")
	}

	return cryptopasta.Encrypt(plaintext, (*[AES256KeyLengthInBytes]byte)(key))
}

// Decrypt decrypts the ciphertext with the provided key and returns the result.
func Decrypt(ciphertext []byte, key *AES256Key) ([]byte, error) {

	if key == nil {
		return nil, errors.New("tried to decrypt with nil key")
	}

	return cryptopasta.Decrypt(ciphertext, (*[AES256KeyLengthInBytes]byte)(key))
}
