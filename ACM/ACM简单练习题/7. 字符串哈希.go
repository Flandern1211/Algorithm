package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(s string, ss *map[string]int) bool {
	if _, exists := (*ss)[s]; !exists {
		(*ss)[s] = 1
		return true
	} else {
		return false
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var N int
	var count int
	fmt.Fscan(in, &N)
	var ss = make(map[string]int, N)

	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(in, &s)
		if check(s, &ss) {
			count++
		}
	}
	fmt.Fprintln(out, count)
}
