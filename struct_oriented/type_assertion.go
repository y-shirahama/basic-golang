package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2) // => 20
	case string:
		fmt.Println(v + "!") // => Mike!
	default:
		fmt.Printf("I don't know %T\n", v) // => I don't know bool
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)
}
