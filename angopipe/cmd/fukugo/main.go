package main

import (
	"angopipe"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// angoとほぼ同じ流れ
// nonceというランダムなバイト列を生成し、暗号化したsourceと共に標準出力として書き出す
func main() {
	gcm, err := angopipe.Prepare()

	if err != nil {
		fmt.Fprintf(os.Stdout, "")
	}

	nonce := make([]byte, 12)
	n, err := io.ReadFull(os.Stdin, nonce)

	if n != 12 || err != nil {
		fmt.Fprintf(os.Stderr, "Can't read nonce: %v\n", err)
		os.Exit(1)
	}

	encrypted, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Decryption Error: %v\n", err)
		os.Exit(1)
	}

	result, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decryption Error: %v\n", err)
		os.Exit(1)
	}
	os.Stdout.Write(result)
}
