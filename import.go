package main

/**
* ライブラリドキュメント：https://golang.org/pkg/
*/
import (
	"fmt"
	"time"
	"os/user"
)

func main() {
	fmt.Println("現在時刻：", time.Now())
	fmt.Println(user.Current())
}