// Package crypto provides a small wrapper around cryptopasta.
package crypto

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/gtank/cryptopasta"
)

// AES256KeyLengthInBytes is the length, in bytes, of a 256-bit AES key
const AES256KeyLengthInBytes = 32

// AES256Key represents a 256-bit AES key
type AES256Key [AES256KeyLengthInBytes]byte

var zeroKey = AES256Key{}

// NewRandomAESKey returns a new, cryptographically generated 256-bit AES key.
// NewRandomAESKey will panic if the source of randomness fails.
func NewRandomAESKey() *AES256Key {
	return (*AES256Key)(cryptopasta.NewEncryptionKey())
}

// NewAESKeyFromBase64 loads the base64-encoded string into the current AES256Key.
func NewAESKeyFromBase64(base64Key string) (*AES256Key, error) {

	if base64Key == "" {
		return nil, errors.New("zero-value base64 string")
	}

	keyBytes, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, err
	}

	if len(keyBytes) != AES256KeyLengthInBytes {
		return nil, fmt.Errorf("expected key length to be %d, was %d", AES256KeyLengthInBytes, len(keyBytes))
	}

	key := &AES256Key{}
	copy(key[:], keyBytes)
	return key, nil
}

// ToBase64 converts the AES key to a base64-encoded string.
func (key *AES256Key) ToBase64() string {

	if key == nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(key[:])
}

// Equal returns a boolean reporting whether a and b are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func Equal(k1, k2 *AES256Key) bool {
	return bytes.Equal(k1[:], k2[:])
}
