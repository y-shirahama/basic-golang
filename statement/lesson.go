package main

import "fmt"

func main() {
	//Q1.以下のスライスから一番小さい数を出力する
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var min int

	for key, num := range l {
		if key == 0 {
			min = num
			continue
		}
		if min >= num {
			min = num
		}
	}
	fmt.Println(min)

	//Q2.以下の果物の価格の合計を出力する
	m := map[string]int{
		"apple": 100,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi": 90,
	}
	sum := 0
	for _, value := range m {
		sum += value
	}
	fmt.Println(sum)
}