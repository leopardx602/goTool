package main

import (
	"fmt"

	"./tool"
)

func main() {
	data, err := tool.OpenJson("read.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
