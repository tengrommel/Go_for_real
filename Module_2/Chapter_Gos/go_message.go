package main

import "time"


func main() {
	messagePrint := func(msg string) {
		println(msg)
	}
	go messagePrint("Hello World")
	go messagePrint("Hello goroutine")
	time.Sleep(time.Second)
}

