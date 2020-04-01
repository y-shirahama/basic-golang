package main

import (
	"fmt"
	"strconv"	//文字列変換ライブラリ
)

func main() {
	//数値→小数点数値
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	//小数点数値→数値
	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	//文字列→数値
	var s string = "14"
	i, _ := strconv.Atoi(s)
	
	//エラーの有無で正常終了するかどうかを判定する
	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println("error!")
	// } else {
	// 	fmt.Printf("%T %v", i, i)
	// }
	fmt.Printf("%T %v\n", i, i)
}