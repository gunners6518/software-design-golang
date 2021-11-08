package main

import (
	"fmt"
	"os"
	"time"
)

// 各ファイルの処理の実行
// 実行前後の時間を計測・表示

func main() {
	logDir, commandName, args := parseArgs()
	stdout, stderr, err := initOut(logDir, commandName)
	if err != nil {
		fmt.Fprintf(stderr, "Can`t create file: %v/n", err)
		os.Exit(1)
	}
	startTime := time.Now()
	ps, err := execution(commandName, args, stdout, stderr)
	if err != nil {
		fmt.Fprintf(stderr, "Can`t exec command %s", commandName)
		os.Exit(1)
	}
	exitTime := time.Now()
	fmt.Println(stdout, ps.String())
	fmt.Fprintf(stdout, "wall clock time=%v system time=%v user time=%v/n", exitTime.Sub(startTime), ps.SystemTime(), ps.UserTime())
}
