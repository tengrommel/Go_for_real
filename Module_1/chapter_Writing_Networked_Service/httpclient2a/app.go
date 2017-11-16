package main

import (
	"net/http"
	"net"
	"time"
	"fmt"
	"io"
	"os"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:true,
			Dial: (&net.Dialer{
				Timeout: 30 * time.Second,
			}).Dial,
		},
	}

	resp, err := client.Get("http://tools.ietf.org/rfc/rfc7540.tx")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	}
