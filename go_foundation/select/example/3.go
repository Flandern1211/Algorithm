package main

import "fmt"

func main() {
	messages := make(chan string, 1) // 有缓冲通道

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message ready, continuing...")
	}

	messages <- "Hello" // 发送消息

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message ready, continuing...")
	}
}

// 输出：
// No message ready, continuing...
// Received message: Hello
