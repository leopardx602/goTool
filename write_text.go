package goTool

import (
	"os"
)

func WriteText(filePath string, data string) error {
	// os.O_WRONLY	只寫
	// os.O_CREATE	建立檔案
	// os.O_RDONLY	只讀
	// os.O_RDWR	讀寫
	// os.O_TRUNC	清空
	// os.O_APPEND	追加
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	// _, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
