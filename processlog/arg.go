package main

import "flag"

// 引数の処理
// return :ログの出力先、実行コマンド名、オプションのリスト
func parseArgs() (string, string, []string) {
	logDir := flag.String(
		"logDir", "", "Log Output directory (default =stderr)")
	flag.Parse()
	return *logDir, flag.Arg(0), flag.Args()[1:]
}
