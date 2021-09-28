package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func writeJson(filePath string, data map[string]interface{}) error {
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

func main() {
	// build data
	tmp := make(map[string]interface{})
	for i := 1; i < 5000; i++ {
		tmp["abc"+strconv.Itoa(i)] = "xyz" + strconv.Itoa(i)
	}
	fmt.Println(tmp)

	// write data
	err := writeJson("file123", tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")

}
