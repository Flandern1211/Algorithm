package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	i := 0
	for {
		select {
		case <-done:
			fmt.Println("Worker 退出")
			return
		case <-time.After(5 * time.Millisecond):
			fmt.Println("Working...", i)
			i++
		}
	}
}

func main() {
	done := make(chan struct{})

	go worker(done)

	time.Sleep(20 * time.Millisecond) // 让 worker 运行一段时间

	close(done) // 发送退出信号

	// 等待 worker 退出，避免 main 函数提前结束
	time.Sleep(10 * time.Millisecond)
}
