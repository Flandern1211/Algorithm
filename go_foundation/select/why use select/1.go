package why_use_select

func main() {
	//假设这里正有一个协程会向dataCh发送数据，但速度不确定
	dataCh := make(chan string)
	//如何实现对dataCh的超时监听？？？
}
