package main

import "fmt"

func add(x int, y int) (int, int) {
	// fmt.Println("add function")	//fmt.Println("add function") -> add function
	// fmt.Println(x + y)	//fmt.Println(x + y) -> 30
	return x + y, x - y
}

func calc(price, item int) (result int) {
	result = price * item
	return result
}

func convert(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1) //fmt.Println(r1) -> 30
	fmt.Println(r2) //fmt.Println(r2) -> -10

	r3 := calc(100, 2)
	fmt.Println(r3) //fmt.Println(r3) -> 200

	f1 := func(x int) {
		fmt.Println("inner func", x)
	}
	f1(1) //-> inner func 1

	func(x int) {
		fmt.Println("inner func", x)
	}(1) //-> inner func 1
}