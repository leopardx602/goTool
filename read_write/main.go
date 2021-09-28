package main

import (
	"fmt"
	"os"
)

func main() {
	// create and delete
	os.Mkdir("folder01", 0777)
	//os.MkdirAll("folder01/test1/test2", 0777)
	//err := os.Remove("folder01")
	err := os.RemoveAll("folder01")
	if err != nil {
		fmt.Println(err)
	}

	//os.Create("test01.txt")

	// read (fast)
	// file, err := os.Open("test01.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()

	// buf := make([]byte, 1024)
	// for {
	// 	n, _ := file.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	os.Stdout.Write(buf[:n])
	// }

	// read (slow, 3 times)
	// fmt.Println(nowTime)
	// inFile, err := ioutil.ReadFile("test01.txt")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// str := string(inFile)
	// fmt.Println(str)

	// write
	// data := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("test01.txt", data, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// write
	// os.O_WRONLY	只寫
	// os.O_CREATE	建立檔案
	// os.O_RDONLY	只讀
	// os.O_RDWR	讀寫
	// os.O_TRUNC	清空
	// os.O_APPEND	追加
	// file, err := os.OpenFile("test01.txt", os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// data := "hello go\n"
	// file.Write([]byte(data))
	// file.WriteString(data)
}
