package goTool

import (
	"encoding/json"
	"os"
)

func ReadJson(filePath string) (map[string]interface{}, error) {
	// open file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// change type
	var data map[string]interface{}
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
