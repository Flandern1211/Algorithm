package chuanzhi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	// 读取第一行：n, m, r1, l2
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	r1, _ := strconv.Atoi(parts[2])
	l2, _ := strconv.Atoi(parts[3])

	// 初始化资源锁定状态，索引从1到n
	isLocked := make([]bool, n+1)
	// 初始未上锁的数量
	count1 := r1         // [1, r1] 未上锁数量
	count2 := n - l2 + 1 // [l2, n] 未上锁数量

	for i := 0; i < m; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())

		if isLocked[x] {
			// 解锁
			isLocked[x] = false
			if x >= 1 && x <= r1 {
				count1++
			}
			if x >= l2 && x <= n {
				count2++
			}
		} else {
			// 上锁
			isLocked[x] = true
			if x >= 1 && x <= r1 {
				count1--
			}
			if x >= l2 && x <= n {
				count2--
			}
		}

		// 输出结果，使用writer提升效率
		fmt.Fprintf(writer, "%d %d\n", count1, count2)
	}
}

func main() {
	// 初始化输入输出，提升大输入场景下的效率
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 读取测试用例数
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		solve(scanner, writer)
	}
}
