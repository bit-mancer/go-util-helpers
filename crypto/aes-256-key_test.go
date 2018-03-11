package crypto_test

import (
	"encoding/base64"

	"github.com/bit-mancer/go-util-helpers/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AES256Key", func() {
	// More of a static assertion, but eh:
	It("is assignable to a [32]byte (as required by crypto functions)", func() {
		var v [32]byte = crypto.AES256Key{}
		v[0] = 0 // prevent unused error
	})

	It("has a factory function that creates a new, random key", func() {
		key1 := crypto.NewRandomAESKey()
		key2 := crypto.NewRandomAESKey()

		Expect(key1).NotTo(Equal(key2))

		emptyKey := &crypto.AES256Key{}

		Expect(key1).NotTo(Equal(emptyKey))
	})

	Describe("NewAESKeyFromBase64", func() {
		It("creates a new AES256Key from a base64 string", func() {
			key, err := crypto.NewAESKeyFromBase64(fixedKeyBase64)
			Expect(err).To(BeNil())
			Expect(key).NotTo(BeNil())
			Expect(key[:]).To(Equal(fixedKey[:]))
		})

		It("requires a valid base64 string", func() {
			key, err := crypto.NewAESKeyFromBase64("")
			Expect(err).NotTo(BeNil())
			Expect(key).To(BeNil())

			key, err = crypto.NewAESKeyFromBase64("bad base64 String")
			Expect(err).NotTo(BeNil())
			Expect(key).To(BeNil())
		})

		It("requires a valid key-length", func() {
			// bad length should fail
			key, err := crypto.NewAESKeyFromBase64(base64.StdEncoding.EncodeToString(fixedKey[:13]))
			Expect(err).NotTo(BeNil())
			Expect(key).To(BeNil())
		})
	})

	Describe("AES256Key.ToBase64", func() {
		It("encodes the key to a base64 string that can be decoded", func() {
			key := fixedKey
			base64String := key.ToBase64()
			Expect(base64String).To(Equal(fixedKeyBase64))

			bytes, err := base64.StdEncoding.DecodeString(base64String)
			Expect(err).To(BeNil())
			Expect(bytes).To(Equal(key[:]))

		})

		It("works on a nil receiver", func() {
			var nilKey *crypto.AES256Key
			s := nilKey.ToBase64()
			Expect(s).To(Equal(""))
		})
	})

	Describe("Equal(k1, k2 *AES256Key)", func() {
		It("compares the equality of two keys", func() {
			key1 := crypto.NewRandomAESKey()
			key2 := crypto.NewRandomAESKey()

			Expect(crypto.Equal(key1, key1)).To(Equal(true))
			Expect(crypto.Equal(key1, key2)).To(Equal(false))

			emptyKey := &crypto.AES256Key{}

			Expect(crypto.Equal(key1, emptyKey)).To(Equal(false))
			Expect(crypto.Equal(emptyKey, emptyKey)).To(Equal(true))
		})
	})
})
