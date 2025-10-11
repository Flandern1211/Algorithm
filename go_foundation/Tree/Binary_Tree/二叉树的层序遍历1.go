package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一:
// 设计思路：
// 层序遍历，就是将每层（1->2->3->...n）按从左往右的顺序来输出
// 而每个节点的左右节点都很好获得，只要从第一个节点，循环获取他的左右节点，然后在按从左往右的顺序来获取左节点的左右节点，右节点的左右节点，依次往下实现即可。
func levelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	//1. 创建一个存放树节点的队列
	// 假设 TreeNode 定义如下，通常二叉树节点包含值、左子节点和右子节点

	q := []*TreeNode{}
	//2. 先将根节点添加进去
	q = append(q, root)
	//3. 直到某个阶段队列中的所有阶点都没有子节点时，这时循环只会将
	//这组节点输出完便没有了
	for len(q) > 0 {
		//每次得到一个将输出的节点
		cur := q[0]
		//更新队列到获取到新的应该输出的节点
		q = q[1:]
		//访问cur节点，输出
		fmt.Println(cur.Val)

		//把cur的左右孩子加入队列
		if cur.Left == nil {
			q = append(q, cur.Left)
		}
		if cur.Right == nil {
			q = append(q, cur.Right)
		}
	}
}

// 方法二:可以输出对应节点层数
func levelOrderTraverse02(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	//记录当前遍历到的层数
	depth := 1

	for len(q) > 0 {
		//获取当前队列长度
		//记住一定要先保存下来,不要使用q.size()作为循环条件，因为在循环过程中队列长度会变化
		sz := len(q)
		for i := 0; i < sz; i++ {
			//弹出队列头
			cur := q[0]
			q = q[1:]
			fmt.Println("depth:=%d,val = %d", depth, cur.Val)

			//加入左右子节点
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
}
