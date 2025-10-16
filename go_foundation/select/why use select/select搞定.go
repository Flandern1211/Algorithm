package why_use_select

import (
	"fmt"
	"time"
)

func main() {
	dataCh := make(chan string)

	// 假设数据在 5ms 或 15ms 时到达
	go func() {
		time.Sleep(5 * time.Millisecond)
		dataCh <- "Success Data Arrived"
	}()

	select {
	case data := <-dataCh:
		// 5ms 时，数据到达。select 立即选择并执行此 case。
		fmt.Println(" 成功收到数据:", data)

	case <-time.After(10 * time.Millisecond):
		// 只有当 10ms 到了，且 dataCh 仍然没有数据时，才会执行此 case。
		fmt.Println("操作超时，未在 10ms 内收到数据。")
	}

	// 程序继续执行，没有 Goroutine 泄露风险！
}
