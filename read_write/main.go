package main

import (
	"fmt"
	"os"

	"github.com/leopardx602/golang/read_write/read"
)

const (
	Filename = "data"
)

func WriteFileAll(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(filename, content string) error {
	// os.O_WRONLY	write only
	// os.O_CREATE	create if not existed
	// os.O_RDONLY	read only
	// os.O_RDWR	read and write only
	// os.O_TRUNC	clear
	// os.O_APPEND	append
	// os.Create(filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) // cover
	if err != nil {
		return err
	}
	defer file.Close()
	// file.Write([]byte(content))
	file.WriteString(content)
	return nil
}

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
