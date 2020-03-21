package main

import "fmt"

func main()  {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)		//fmt.Println(n) -> [1 2 3 4 5 6]
	fmt.Println(n[2])	//fmt.Println(n[2]) -> 3
	fmt.Println(n[2:4])	//fmt.Println(n[2:4]) -> [3 4]
	fmt.Println(n[:2]) 	//fmt.Println(n[:2]) -> [1 2]
	fmt.Println(n[2:])	//fmt.Println(n[2:]) -> [3 4 5 6]
	fmt.Println(n[:])		//fmt.Println(n[:]) -> [1 2 3 4 5 6]

	n[2] = 100
	fmt.Println(n)		//fmt.Println(n) -> [1 2 100 4 5 6]

	var board = [][]int {
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)	//fmt.Println(board) -> [[0 1 2] [3 4 5] [6 7 8]]

	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)		//fmt.Println(n) -> [1 2 100 4 5 6 100 200 300 400]
}