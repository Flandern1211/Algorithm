package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	fmt.Scan(&a, &b)
	if a == b {
		fmt.Println(0)
		return
	}
	if a == -b {
		fmt.Println(3)
		return
	}
	if a == 0 && b != 0 {
		fmt.Println(2)
		return
	}
	if b == 0 && a != 0 {
		fmt.Println(1)
		return
	}
	fmt.Println(-1)

}
