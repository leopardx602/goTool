package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	exePath := filepath.Dir(ex)
	fmt.Println("exePath:", exePath)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println("absPath:", dir)

	dir2, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("workingDirPath:", dir2)
}
