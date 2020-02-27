package main

import (
	"fmt"
)

// varで宣言した変数は、関数外でも利用できる
var (
	num int = 2
	float float64 = 2.4
	str string = "Go言語2"
	on, off bool = true, false
)

func foo_1() {
	var i int = 1
	var f64 float64 = 1.2
	var s string = "Go言語"
	var t, f bool = true, false

	fmt.Println(i, f64, s, t, f)
}

func foo_2() {
	xi := 1
	xf64 := 1.2
	xs := "Go言語"
	xt, xf := true, false

	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
}

func main() {
	fmt.Println(num, float, str, on, off)
	foo_1()
	foo_2()
}