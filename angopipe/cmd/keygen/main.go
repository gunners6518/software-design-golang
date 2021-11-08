package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// 32バイトの乱数から暗号にする文字列を生成
func main() {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	readableKey := base64.StdEncoding.EncodeToString(key)
	fmt.Println(readableKey)
}
