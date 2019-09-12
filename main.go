package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	prefix          *string
	targetExtension *string
	dir             *string
	N               int
)

func main() {
	//
	// コマンドライン引数をセットする
	// ex: ./main.go -p=abc- -t=.txt

	//　接頭辞
	prefix = flag.String("p", "Hello world!", "option test")
	// 対象のディレクトリ
	dir = flag.String("d", "testdata/testDir01/", "set directory")
	// 変更を加えるファイルの拡張子
	targetExtension = flag.String("t", ".txt", "set target extension")
	// 消去する文字数
	flag.IntVar(&N, "n", 1, "set delete num")

	flag.Parse()

	// コマンド
	cmd := flag.Arg(0)
	fmt.Println(cmd)

	// todo: addPrefix と delName 両方の機能を呼び出せるようにする

	// コマンド引数が 'del' の時はdelNameを実行する
	if cmd == "del" {
		err := filepath.Walk(*dir, delName)
		if err != nil {
			fmt.Println(err)
		}
	}

	// コマンド引数が存在しない場合は、addPrefixを実行する
	if cmd == "" {
		err := filepath.Walk(*dir, addPrefix)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// ファイル名から指定した数だけ文字数を削除する
func delName(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if filepath.Ext(path) == *targetExtension {
		fmt.Println(path)

		// ファイル名を取得
		oldName := filepath.Base(path)

		// ディレクトリ名を取得
		dirName := filepath.Dir(path)

		// ファイル名を任意の文字だけ削除
		n := N
		newName := oldName[n:]

		newPath := filepath.Join(dirName + "/" + newName)

		errRename := os.Rename(path, newPath)
		if errRename != nil {
			fmt.Println(errRename)
		}
		//fmt.Printf("dirName: %v, oldName: %v, newName: %v\n", dirName, oldName, newName)
	}

	return nil
}

// ファイル名に接頭辞を付与する
func addPrefix(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// ルートディレクトリも path対象となってしまうため、対象ファイルの拡張子を指定する

	if filepath.Ext(path) == *targetExtension {
		fmt.Println(path)

		// ファイル名を取得
		oldName := filepath.Base(path)

		// ディレクトリ名を取得
		dirName := filepath.Dir(path)

		// 接頭辞を付与する
		newName := *prefix + oldName

		// ファイル名を変更
		newPath := filepath.Join(dirName + "/" + newName)

		errRename := os.Rename(path, newPath)
		if errRename != nil {
			fmt.Println(errRename)
		}
		//fmt.Printf("dirName: %v, oldName: %v, newName: %v\n", dirName, oldName, newName)
	}

	return nil
}
