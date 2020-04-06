package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	fmt.Println(n)	// => 100
	fmt.Println(&n) // => 0xc00001a090(メモリ上に割り振られたアドレス値)

	var p *int = &n
	fmt.Println(p)	// => 0xc00001a090(メモリ上に割り振られたアドレス値)
	fmt.Println(*p)	// => 100(アドレスが指しているメモリ値)

	one(&n)
	fmt.Println(n)	// => 1
	//fmt.Println(&*&n)
}