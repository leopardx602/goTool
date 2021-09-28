// slow 5~10 times
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func openJson2(filePath string) (map[string]interface{}, error) {
	// open file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// change type
	var data map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	data, err := openJson2("read.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
