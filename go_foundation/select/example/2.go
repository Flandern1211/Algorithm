package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// 模拟一个慢速操作
	go func() {
		time.Sleep(2 * time.Second) // 超过 500ms
		ch <- "Message from worker"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(500 * time.Millisecond): // 超时通道
		fmt.Println("Timeout! Operation took too long.")
	}
}

// 输出：
// Timeout! Operation took too long.
