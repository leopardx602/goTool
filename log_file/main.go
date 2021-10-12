package main

import (
	"log"
	"os"
)

func main() {
	// open log file
	logPath := "log/logfile.log"
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	defer f.Close()

	// method1
	log.SetOutput(f)
	log.Println("Log test")

	// method2, definition by self
	const (
		Ldate         = 1 << iota     // the date: 2009/01/23
		Ltime                         // the time: 01:23:23
		Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
		Llongfile                     // full file name and line number: /a/b/c/d.go:23
		Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
		LstdFlags     = Ldate | Ltime // initial values for the standard logger
	)
	logger := log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	logger.SetOutput(f)
	logger.Println("Log test2")
}
