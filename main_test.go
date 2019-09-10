package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// rename関数により、ファイル名が変更されることを確認

func TestRename(t *testing.T) {
	err := filepath.Walk("testdata/testDir01/", Rename)
	if err != nil {
		fmt.Println(err)
	}

	// ディレクトリ内のファイル全てに対して、名前が変更されたか確認する
	err2 := filepath.Walk("testdata/testDir01/", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			// ファイル名を取得
			fileName := filepath.Base(path)

			// 予想ファイル名との一致を確認 "new-"で始まっていたらOKとする
			if fileName[:4] != "new-" {
				t.Fatalf("want: %v, got: %v", "new-xx.jpg", fileName)
			}
		}
		return nil
	})

	if err2 != nil {
		fmt.Println(err2)
	}

}
