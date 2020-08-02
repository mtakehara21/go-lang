package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/**
作成するcatコマンド􏰀仕様
● 引数でファイルパス􏰀一覧を貰い、そ􏰀ファイルを与えられた順に標準出
力にそ􏰀まま出力するコマンドを作ってください。
● また、-nオプションを指定すると、行番号を各行につけて表示されるように してください。
● なお、行番号􏰁すべて􏰀ファイルで通し番号にしてください。
**/

func main() {
	// オプションを設定
	var n = flag.Bool("n", false, "通し番号を付与する")
	flag.Parse()

	// 引数を取得
	var files = flag.Args()
	// 実行ファイルパスの取得
	path, err := os.Executable()
	if err != nil {
		print("err")
	}
	path = filepath.Dir(path)

	for i := 0; i < len(files); i++ {
		// ファイルの展開
		sf, err := os.Open(filepath.Join(path, files[i]))
		if err != nil {
			print("err")
		} else {
			x := 1
			scanner := bufio.NewScanner(sf)
			// １行ずつ取得し出力
			for scanner.Scan() {
				if *n {
					fmt.Printf("%v: ", x)
				}
				x++
				fmt.Println(scanner.Text())
			}
		}
	}
}
