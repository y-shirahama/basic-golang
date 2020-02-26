package main

import "fmt"

func init() {
	fmt.Println("Init!")
}

func main() {
	bazz()
	fmt.Println("Hello World!", "Go言語で書いてみた")
}

func bazz() {
	fmt.Println("Bazz")
}
