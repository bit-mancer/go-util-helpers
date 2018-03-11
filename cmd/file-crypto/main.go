/*
Simple file encryption/decrypt.
The file must fit in available memory.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bit-mancer/go-util-helpers/crypto"
)

var (
	encrypt    bool
	decrypt    bool
	base64Key  string
	inputFile  string
	outputFile string
)

func init() {
	flag.BoolVar(&encrypt, "e", false, "Encrypt.")
	flag.BoolVar(&decrypt, "d", false, "Decrypt.")
	flag.StringVar(&base64Key, "k", "", "Base64-encoded AES-256 key.")
	flag.StringVar(&inputFile, "i", "", "Input file; if not provided, input will be read from stdin.")
	flag.StringVar(&outputFile, "o", "", "Output file; if not provided, output will be sent to stdout.")
}

var (
	fileMode os.FileMode = 0644
)

func readInput() ([]byte, error) {
	if inputFile == "" {
		return ioutil.ReadAll(os.Stdin)
	}

	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Error getting information on the input file: %v", err)
	}

	fileMode = fileInfo.Mode()

	return ioutil.ReadFile(inputFile)
}

func writeOutput(data []byte) error {
	if outputFile != "" {
		return ioutil.WriteFile(outputFile, data, fileMode)
	}

	_, err := os.Stdout.Write(data)
	return err
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-e | -d] -k <key> [-i <input-file>] [-o <output-file>]\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()

	switch {
	case base64Key == "":
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

	data, err := readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var renderedBytes []byte

	if encrypt {
		renderedBytes, err = crypto.Encrypt(data, key)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error encrypting:", err)
			os.Exit(1)
		}
	} else if decrypt {
		renderedBytes, err = crypto.Decrypt(data, key)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error decrypting:", err)
			os.Exit(1)
		}
	} else {
		panic("no mode specified")
	}

	if err := writeOutput(renderedBytes); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
