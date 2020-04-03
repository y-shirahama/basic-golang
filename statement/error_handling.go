package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./statement/if.go")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln(err.Error())
	}
}