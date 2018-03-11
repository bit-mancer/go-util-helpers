package crypto_test

import (
	"encoding/base64"

	"github.com/bit-mancer/go-util-helpers/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncryptStringToBase64", func() {
	It("encrypts a string", func() {
		plaintext := "test"
		ciphertext, err := crypto.EncryptStringToBase64(plaintext, &fixedKey)
		Expect(ciphertext).NotTo(Equal(""))
		Expect(err).To(BeNil())

		Expect(ciphertext).NotTo(Equal(plaintext))
	})

	It("returns a base64-encoded string", func() {
		ciphertext, err := crypto.EncryptStringToBase64("test", &fixedKey)
		Expect(ciphertext).NotTo(Equal(""))
		Expect(err).To(BeNil())

		// this is a bit whitebox-ish because we express the padding used (standard)
		_, err = base64.StdEncoding.DecodeString(ciphertext)
		Expect(err).To(BeNil())
	})

	It("requires a valid key", func() {
		ciphertext, err := crypto.EncryptStringToBase64("test", nil)
		Expect(ciphertext).To(Equal(""))
		Expect(err).NotTo(BeNil())
	})
})

var _ = Describe("DecryptStringFromBase64", func() {
	It("returns plaintext originally encrypted by EncryptStringToBase64 with the same key", func() {
		plaintext := "test"
		ciphertext, err := crypto.EncryptStringToBase64(plaintext, &fixedKey)
		Expect(ciphertext).NotTo(Equal(""))
		Expect(err).To(BeNil())

		plaintext2, err := crypto.DecryptStringFromBase64(ciphertext, &fixedKey)
		Expect(plaintext2).To(Equal(plaintext))
		Expect(err).To(BeNil())

		plaintext, err = crypto.DecryptStringFromBase64(ciphertext, crypto.NewRandomAESKey())
		Expect(plaintext).To(Equal(""))
		Expect(err).NotTo(BeNil())
	})

	It("requires a valid key", func() {
		plaintext, err := crypto.DecryptStringFromBase64(fixedKeyBase64, nil)
		Expect(plaintext).To(Equal(""))
		Expect(err).NotTo(BeNil())
	})
})
