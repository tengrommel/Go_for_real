package main

import (
	"net/http"
	"time"
	"fmt"
	"io"
	"os"
)

func main() {
	client := &http.Client{
		Timeout: 21 * time.Second,
	}
	resp, err := client.Get("http://tools.ietf.org/rfc/rfc7540.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	}