package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine 1: 立即向 c1 发送数据
	go func() {
		c1 <- "来自 c1 的消息"
	}()

	// Goroutine 2: 立即向 c2 发送数据
	go func() {
		c2 <- "来自 c2 的消息"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Case 1 获胜:", msg1)
	case msg2 := <-c2:
		fmt.Println("Case 2 获胜:", msg2)
	}
}
