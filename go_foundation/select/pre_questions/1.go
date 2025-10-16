package main

import "fmt"

func main() {
	c := make(chan int)

	select {
	case <-c:
		fmt.Println("read")
	case c <- 1:
		fmt.Println("write")
	}
}
