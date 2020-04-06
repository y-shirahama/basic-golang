package main

import "fmt"

/**
 * 構造体
 */
type Vertex struct {
	X, Y int
	Str string
}

/**
 * 構造体の値渡し
 */
func changeVertex(v Vertex) {
	v.X = 1000
}

/**
 * ポインタ型の構造体の値渡し
 */
func changeVertex2(v *Vertex) {
	v.X = 1000
	//(*v).X = 1000
}

func main()  {
	v := Vertex{1, 2, "golang"}
	changeVertex(v)
	fmt.Println(v)								// => {1 2 golang}

	v2 := &Vertex{1, 2, "golang"}
	changeVertex2(v2)
	fmt.Println(*v2)							// => {1000 2 golang}
}