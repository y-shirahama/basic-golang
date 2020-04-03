package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	file, _ := os.OpenFile(logFile, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("test.log")
	_, err := os.Open("test.go")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v", "python", "kotlin")

	log.Fatalf("%T %v", "go", "php")
	log.Fatalln("error!")

	fmt.Println("ok?")
}