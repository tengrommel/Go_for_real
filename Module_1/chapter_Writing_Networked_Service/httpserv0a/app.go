package main

import "fmt"
import "net/http"

type msg string

func (m msg)ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	res.Header().Add("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, m)
}

func main() {
	msgHandler := msg("Hello from high above!")
	server := http.Server{Addr:":4040", Handler: msgHandler}
	server.ListenAndServe()
}
