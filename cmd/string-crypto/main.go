package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bit-mancer/go-util-helpers/crypto"
)

var (
	encrypt   bool
	decrypt   bool
	base64Key string
)

func init() {
	flag.BoolVar(&encrypt, "e", false, "Encrypt.")
	flag.BoolVar(&decrypt, "d", false, "Decrypt.")
	flag.StringVar(&base64Key, "k", "", "Base64-encoded AES-256 key.")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-e | -d] -k <key> <text>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
	}

	text := flag.Arg(0)

	switch {
	case base64Key == "", text == "":
		fallthrough
	case encrypt && decrypt:
		fallthrough
	case !encrypt && !decrypt:
		flag.Usage()
	}

	key, err := crypto.NewAESKeyFromBase64(base64Key)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading the base64-encoded AES-256 key:", err)
		os.Exit(1)
	}

	var renderedText string

	if encrypt {
		if renderedText, err = crypto.EncryptStringToBase64(text, key); err != nil {
			fmt.Fprintln(os.Stderr, "Error encrypting:", err)
			os.Exit(1)
		}
	} else if decrypt {
		if renderedText, err = crypto.DecryptStringFromBase64(text, key); err != nil {
			fmt.Fprintln(os.Stderr, "Error decrypting:", err)
			os.Exit(1)
		}
	} else {
		panic("no mode specified")
	}

	fmt.Println(renderedText)
}
