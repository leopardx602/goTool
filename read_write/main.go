package main

import (
	"fmt"

	"github.com/leopardx602/golang/read_write/read"
)

const (
	Filename = "data"
)

func main() {
	// content, err := read.ReadFileAll(Filename)
	// if err != nil {
	// 	fmt.Println("err")
	// }
	// fmt.Println(content)

	// content, err := read.ReadFilePointer(Filename)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(content)

	ch := make(chan []byte, 100)
	go func() {
		defer close(ch)
		if err := read.ReadLineToChannel(Filename, ch); err != nil {
			panic(err)
		}
	}()
	for val := range ch {
		fmt.Println(string(val))
	}

	// if err := WriteFileAll(Filename, "444\n222\n333"); err != nil {
	// 	panic(err)
	// }

	// if err := WriteFile(Filename, "111\n222\n333"); err != nil {
	// 	panic(err)
	// }

	// // create directory
	// if err := os.Mkdir("folder01", 0777); err != nil {
	// 	panic(err)
	// }
	// if err := os.MkdirAll("folder01/test1/test2/test3", 0777); err != nil {
	// 	panic(err)
	// }

	// // remove directory
	// if err := os.Remove("folder01"); err != nil {
	// 	panic(err)
	// }
	// if err := os.RemoveAll("folder01"); err != nil {
	// 	panic(err)
	// }
}
