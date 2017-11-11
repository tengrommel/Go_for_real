package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	makeEvenNums(4, ch)
	for i:=0;i<=3;i++  {
		fmt.Println(<-ch)
	}
}

func makeEvenNums(count int, in chan<- int)  {
	for i:=0;i<count;i++ {
		in <- 2*i
	}
}
