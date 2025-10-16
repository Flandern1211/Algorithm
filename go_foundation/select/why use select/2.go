package why_use_select

import (
	"fmt"
	"time"
)

func main() {
	// 1. 定义两个 Channel
	dataCh := make(chan string)   // 接收外部数据
	resultCh := make(chan string) // 统一的结果通道
	//数据接收goroutine
	go func() {
		x := <-dataCh
		resultCh <- x
	}()

	//超时处理goroutine
	go func() {
		// 阻塞 10ms，等待超时信号
		<-time.After(10 * time.Second)

		// 转发超时结果
		resultCh <- "接收超时,不再接收"
	}()
	firstResult := <-resultCh
	fmt.Println(firstResult)
}
