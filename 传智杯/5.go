package chuanzhi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	//存储常数的切片
	var terms = make([]int, 0)
	var signed int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if i+3 >= len(s) { // 边界检查
				break
			}
			if s[i+2] == '+' {
				signed = 1
			} else {
				signed = -1
			}
			num, _ := strconv.Atoi(string(s[i+3]))
			terms = append(terms, num*signed)
			i += 4
		} else {
			i++
		}
	}

	if len(terms) == 0 {
		fmt.Print(0)
		return
	}

	currentA := 1
	currentB := terms[0]

	for _, d := range terms[1:] {
		newA := (currentA*d + 1*currentB)
		newB := (currentB * d)

		currentA = newA
		currentB = newB
	}
	fmt.Fprintln(out, currentA%10074)
}
