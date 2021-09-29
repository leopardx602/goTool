package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// method1
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}

	// method2
	files2, _ := filepath.Glob("*.go")
	fmt.Println(files2)
	for _, f := range files2 {
		fmt.Println(f)
	}
}
