package main

import (
	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost/"
	ws, err := websocket.Dial("ws://localhost:4000/websocket", "", origin)
	if err != nil{
		panic(err)
	}

	_, err = ws.Write([]byte("Hello?Anybody here?"))
	if err != nil{
		panic(err)
	}

	msg := make([]byte, 512)
	_, err = ws.Read(msg)

}
