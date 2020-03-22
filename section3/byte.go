package main

import "fmt"

func main() {
	b := []byte{72, 73}
	fmt.Println(b) //fmt.Println(b) -> [72 73]

	//https://www.ascii-code.com -ASCII Code - The extended ASCII table-
	fmt.Println(string(b)) //fmt.Println(string(b)) -> HI

	c := []byte("HI")
	fmt.Println(c)         //fmt.Println(c) -> [72 73]
	fmt.Println(string(c)) //fmt.Println(string(c)) -> HI
}
