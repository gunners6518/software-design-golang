package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// 出力先の生成
// -logDirオプションなし:標準出力(stdout),標準エラー(stderr)を返す
// -logDirオプションあり:ファイル生成

func initOut(logDir, commandName string) (stdout, stderr io.Writer, err error) {
	if logDir == "" {
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		ts := time.Now().Unix()

		stdoutFileName := fmt.Sprintf("%s-%v-stdout.log", commandName, ts)
		stderrFile, err := os.Create(filepath.Join(logDir, stdoutFileName))

		if err != nil {
			return nil, nil, err
		}

		stderr = io.MultiWriter(os.Stderr, stderrFile)
	}
	return
}
