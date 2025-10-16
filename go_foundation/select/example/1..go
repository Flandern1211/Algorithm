package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Result from Server 1"
}

func server2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Result from Server 2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go server1(c1)
	go server2(c2)

	// 同时等待 c1 和 c2，哪个先准备好就执行哪个 case
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// 输出可能是：
// Result from Server 1
// Result from Server 2
