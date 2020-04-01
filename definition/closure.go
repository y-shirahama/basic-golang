package main

import "fmt"

func incrementGenerator() (func () int) {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func (radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	counter := incrementGenerator()
	fmt.Println(counter()) //fmt.Println(counter()) -> 1
	fmt.Println(counter()) //fmt.Println(counter()) -> 2
	fmt.Println(counter()) //fmt.Println(counter()) -> 3

	c1 := circleArea(3.14)
	fmt.Println(c1(2)) //fmt.Println(c1(2)) -> 12.56
	fmt.Println(c1(3)) //fmt.Println(c1(3)) -> 28.259999999999998

	c2 := circleArea(3)
	fmt.Println(c2(2)) //fmt.Println(c2(2)) -> 12
	fmt.Println(c2(3)) //fmt.Println(c2(3)) -> 18
}