package main

import (
	"bufio"
	"fmt"
	"os"
)

var stack = make([]int, 0)
var p = &stack
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

// 将整数x入栈
func Push(x int) {
	(*p) = append(*p, x)
}

// 若栈非空，则删除栈顶元素；否则输出
func Pop() {
	if len(*p) != 0 {
		(*p) = (*p)[:len(*p)-1]
	} else {
		fmt.Fprintln(out, "Empty")
	}
}

// 若栈非空，则输出栈顶元素；否则输出 Empty
func Query() {
	if len(*p) != 0 {
		fmt.Fprintln(out, (*p)[len(*p)-1])
	} else {
		fmt.Fprintln(out, "Empty")
	}
}

func Size() {
	fmt.Fprintln(out, len(*p))
}

func main() {
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var op string
		fmt.Fscan(in, &op)
		switch op {
		case "push":
			var num int
			fmt.Fscan(in, &num)
			Push(num)
		case "pop":
			Pop()
		case "query":
			Query()
		case "size":
			Size()
		}
	}

}
