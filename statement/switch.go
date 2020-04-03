package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Windows!")
	default:
		fmt.Println("Default!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")	
	}
}