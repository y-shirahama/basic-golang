package main

import "fmt"

func thirdPartyConnectDB() {
	//Go言語ではpanic()を使うことをあまり推奨していない
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		//recover()はpanic()で起きた処理をキャッチする
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
