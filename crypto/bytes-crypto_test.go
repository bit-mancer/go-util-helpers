package crypto_test

import (
	"bytes"

	"github.com/bit-mancer/go-util-helpers/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encrypt", func() {
	It("encrypts a byte-slice", func() {
		plaintext := []byte("test")
		ciphertext, err := crypto.Encrypt(plaintext, &fixedKey)
		Expect(ciphertext).NotTo(Equal([]byte("")))
		Expect(err).To(BeNil())

		Expect(bytes.Equal(ciphertext, plaintext)).To(Equal(false))
	})

	It("requires a valid key", func() {
		ciphertext, err := crypto.Encrypt([]byte("test"), nil)
		Expect(bytes.Equal(ciphertext, []byte(""))).To(Equal(true))
		Expect(err).NotTo(BeNil())
	})
})

var _ = Describe("Decrypt", func() {
	It("returns plaintext originally encrypted by Encrypt with the same key", func() {
		plaintext := []byte("test")
		ciphertext, err := crypto.Encrypt(plaintext, &fixedKey)
		Expect(bytes.Equal(ciphertext, []byte(""))).To(Equal(false))
		Expect(err).To(BeNil())

		plaintext2, err := crypto.Decrypt(ciphertext, &fixedKey)
		Expect(bytes.Equal(plaintext2, plaintext)).To(Equal(true))
		Expect(err).To(BeNil())

		plaintext2, err = crypto.Decrypt(ciphertext, crypto.NewRandomAESKey())
		Expect(bytes.Equal(plaintext2, []byte(""))).To(Equal(true))
		Expect(err).NotTo(BeNil())
	})

	It("requires a valid key", func() {
		plaintext, err := crypto.Decrypt([]byte("test"), nil)
		Expect(bytes.Equal(plaintext, []byte(""))).To(Equal(true))
		Expect(err).NotTo(BeNil())
	})
})
