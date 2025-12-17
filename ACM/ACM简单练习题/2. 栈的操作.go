package main

import (
	"bufio"
	"fmt"
	"os"
)

var stack = make([]int, 0)
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

// 将整数x入栈
func Push(x int) {
	if len(stack) == 0 {
		return
	}
	stack[len(stack)-1] = x
}

func Pop() {
	if len(stack) != 0 {
		stack = stack[:len(stack)-1]
	} else {
		fmt.Fprintln(out, "Empty")
	}
}

func Query() {
	if len(stack) != 0 {
		stack = stack[:len(stack)-1]
		fmt.Fprintln(out, stack[len(stack)-1])
	} else {
		fmt.Fprintln(out, "Empty")
	}
}

func Size() {
	fmt.Fprintln(out, len(stack))
}

func main() {

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
