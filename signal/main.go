package main

import (
	"log"
	"time"
	"os"
	"syscall"
	"os/signal"
)

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		log.Println("done")
	}()

	timer := time.NewTimer(time.Second * 5)

	select {
	case <- timer.C:
	case <- sigCh:
	}
}
