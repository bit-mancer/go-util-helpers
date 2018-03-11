package crypto_test

import (
	"github.com/bit-mancer/go-util-helpers/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var fixedKey = crypto.AES256Key{144, 219, 85, 9, 191, 207, 189, 121, 11, 162, 77, 82, 15, 129, 176, 78, 102, 55, 120, 13, 173, 171, 130, 177, 247, 157, 53, 121, 113, 69, 189, 25}

const fixedKeyBase64 = "kNtVCb/PvXkLok1SD4GwTmY3eA2tq4Kx9501eXFFvRk="

func TestCrypto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crypto Suite")
}
