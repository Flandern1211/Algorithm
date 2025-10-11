package Tree

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	//1.创建一个切片来存储结果
	result := []int{}
	//
	for _, children := range root.Children {
		temp := postorder(children)
		result = append(result, temp...)
	}

	result = append(result, root.Val)
	return result
}
