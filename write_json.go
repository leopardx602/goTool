package goTool

import (
	"encoding/json"
	"os"
)

func WriteJson(filePath string, data map[string]interface{}) error {
	// open file
	file, err := os.OpenFile("write.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write data
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	file.Write(dataByte)
	return nil
}
