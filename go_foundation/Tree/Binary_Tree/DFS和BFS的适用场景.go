package main

//方法一:DFS（递归来实现）
//实现要点:
import (
	"math"
)

func minDepth(root *TreeNode) int {
	//
	if root == nil {
		return 0
	}
	//记录最小深度
	minDepthValue := math.MaxInt32

	//记录当前遍历到的节点的深度
	currentDepth := 0

	//匿名函数，通过递归来得到最小深度
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		//3.1前序位置进入时增加当前深度，前序位置进入（访问节点前）的相当于直接对本次节点进行操作，而还没有进入到它的左右子节点，得到的就是该节点的深度。
		currentDepth++

		//3.2判断当前节点是否为叶子节点，是的话进行判断并更新最小深度
		if root.Left == nil && root.Right == nil {
			minDepthValue = min(minDepthValue, currentDepth)

		}

		traverse(root.Left)
		traverse(root.Right)
		//3.3当一个分支的最小深度判断更新完后就会迂回到它们的上层节点，也就是tracerse（root.Right）执行完的后序位置的，这时的currentDepth因为刚从下层回来，还是下层的深度，所以需要-1来更新实时深度。
		currentDepth--
	}

	traverse(root)

	return minDepthValue
}

// 方法二：BFS
// 设计:
func minDepth02(root *TreeNode) int {
	if root == nil {
		return 0
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
			if cur.Left == nil && cur.Right == nil {
				return depth
			}

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
	return depth
}
