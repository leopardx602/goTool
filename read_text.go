package goTool

import (
	"os"
)

func ReadText(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var data string
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		//os.Stdout.Write(buf[:n])
		data += string(buf[:n])
	}
	return data, nil
}
