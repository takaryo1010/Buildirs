package buildirs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Buildirs(path string, NewFileName []string, Nums []int, n int) error {
	fileName := NewFileName[0] // 変更: NewFileName[0]を別の変数に代入する
	Name_Separated := strings.Split(fileName, "*")

	if len(NewFileName) == 1 && len(Name_Separated) == 1 {
		err := os.MkdirAll(NewFileName[0], 0750)
		if err != nil {
			return err
		}
		return nil

	}

	if len(Name_Separated) == 1 {
		fileName = fileName + "\\" + NewFileName[1]                    // 変更: ファイル名を更新する
		NextFileName := append([]string{fileName}, NewFileName[2:]...) // 変更: 更新されたファイル名を使用して、配列を再構築する
		Buildirs(path, NextFileName, Nums, n)
		return nil
	}
	n++
	for i := 0; i < Nums[n]; i++ {
		
		formattedPath := filepath.Join(path, fmt.Sprintf("%v%d%v", Name_Separated[0], i+1, Name_Separated[1]))
		if len(NewFileName) > 1 { // 変更: NewFileNameの要素数が1より大きい場合にのみ、再帰呼び出しを行う
			newFileName := filepath.Join(formattedPath, NewFileName[1])       // 変更: 更新されたフォルダ名と元のファイル名を使用して、新しいファイルパスを作成する
			NextFileName := append([]string{newFileName}, NewFileName[2:]...) // 変更: 新しいファイルパスを使用して、配列を再構築する
			Buildirs(path, NextFileName, Nums, n)
		}

		err := os.MkdirAll(formattedPath, 0750)
		if err != nil {
			return err
		}
	}

	return nil
}
