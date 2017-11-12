package main

import (
	"fmt"
	"sync"
)

const MAX = 1000

func main() {
	values:=make(chan int, MAX)
	result:=make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i:=0;i<MAX;i++{
			if(i%3)==0 || (i%5) == 0 {
				values <- i
			}
		}
		close(values)
	}()

	work := func() {
		defer wg.Done()
		r := 0
		for i:=range values {
			r+=i
		}
		result <- r
	}

	go work()
	go work()
	wg.Wait()
	total := <-result + <-result
	fmt.Println("Total:", total)
}
