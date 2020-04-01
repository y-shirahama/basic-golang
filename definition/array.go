package main

import (
	"fmt"
)

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
} 