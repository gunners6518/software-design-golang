package main

import (
	"io"
	"os"
	"os/exec"
)

// 子プロセスの起動

func execution(commandName string, args []string, stdout, stderr io.Writer) (*os.ProcessState, error) {
	cmd := exec.Command(commandName, args...)

	// 子プロセスの出力・エラーを取得
	childStdout, _ := cmd.StdoutPipe()
	childStderr, _ := cmd.StderrPipe()

	//　標準出力・エラーに子プロセスの標準出力・エラーを入れる
	go io.Copy(stdout, childStdout)
	go io.Copy(stderr, childStderr)
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return cmd.ProcessState, nil
}
