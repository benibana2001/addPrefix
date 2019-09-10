package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("testdata/testDir01/", Rename)
	if err != nil {
		fmt.Println(err)
	}
}

func Rename(path string, info os.FileInfo, err error) error {
	if err != nil {
			return err
		}

		// ルートディレクトリも path対象となってしまうため、対象ファイルの拡張子を指定する

		if filepath.Ext(path) == ".txt" {
			fmt.Println(path)

			// ファイル名を取得
			oldName := filepath.Base(path)

			// ディレクトリ名を取得
			dirName := filepath.Dir(path)

			// 接頭辞を付与する
			newName := "new-" + oldName

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

