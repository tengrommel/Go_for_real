package main

import(
	"net/http"
	"golang.org/x/net/websocket"
)

func Handle(ws *websocket.Conn)  {

	ws.Write([]byte("I am alive"))
}

func main() {
	http.Handle("/websocket", websocket.Handler(Handle))
	err:= http.ListenAndServe(":4000",nil)
	if err != nil{
		panic(err)
	}
}
