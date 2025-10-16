package main

import (
	"fmt"
	"time"
)

func worker(workCh <-chan string, quitCh <-chan bool) {
	for {
		select {
		case work := <-workCh:
			// 处理工作
			fmt.Println("Processing work:", work)
		case <-quitCh:
			// 收到退出信号，退出循环
			fmt.Println("Worker received quit signal, shutting down.")
			return
		}
	}
}

func main() {
	workCh := make(chan string)
	quitCh := make(chan bool)

	go worker(workCh, quitCh)

	// 发送一些工作
	workCh <- "Task A"
	time.Sleep(100 * time.Millisecond)
	workCh <- "Task B"
	time.Sleep(100 * time.Millisecond)

	// 发送退出信号
	quitCh <- true
	time.Sleep(100 * time.Millisecond) // 等待 worker 退出
	fmt.Println("Main routine finished.")
}
