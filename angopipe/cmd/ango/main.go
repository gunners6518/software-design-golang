package main

import (
	"angopipe"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// nonceというランダムなバイト列を生成し、暗号化したsourceと共に標準出力として書き出す
func main() {
	// prepareを実行して初期設定
	gcm, err := angopipe.Prepare()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 暗号化したデータをsourceとして受け取る
	source, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encription Error: %v\n", err)
		os.Exit(1)
	}

	// ランダムなバイト列を作る
	nonce := make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)
	os.Stdout.Write(nonce)

	// 標準出力として書き出す
	result := gcm.Seal(nil, nonce, source, nil)
	os.Stdout.Write(result)

}
