package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)          // fmt.Println(m) -> map[apple:100 banana:200]
	fmt.Println(m["apple"]) // fmt.Println(m["apple"]) -> 100
	m["banana"] = 300
	fmt.Println(m) // fmt.Println(m) -> map[apple:100 banana:300]
	m["new"] = 500
	fmt.Println(m) // fmt.Println(m) -> map[apple:100 banana:300 new:500]

	fmt.Println(m["nothing"]) // fmt.Println(m["nothing"]) -> 0

	v, ok := m["apple"]
	fmt.Println(v, ok) // fmt.Println(v, ok) -> 100 true

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2) // fmt.Println(v2, ok2) -> 0 false

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2) // fmt.Println(m2) -> map["pc":5000]


	/** メモリ状に入れるmapがないため、エラーが発生する。mapは宣言だけでは初期化されない
	var m3 map[string]int　// panic: assignment to entry in nil map
	m3["pc"] = 5000
	fmt.Println(m3)	
	*/

	/* マップ型もスライスも宣言時のゼロ値はnilとなる */
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
