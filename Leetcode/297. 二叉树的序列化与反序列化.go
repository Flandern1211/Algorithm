package Leetcode

import (
	"strconv"
	"strings"
)

type Codec struct {
	SEP  string
	NULL string
}

func Constructor() Codec {
	return Codec{
		SEP:  ",",
		NULL: "#",
	}
}

func (this Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	this._serialize(root, &sb)
	return sb.String()
}

// 不需要返回string,因为传入&sb,在递归的同时sb的值就已经在变了
func (this Codec) _serialize(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(this.NULL)
		sb.WriteString(this.SEP)
	} else {
		//前序位置操作
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteString(this.SEP)
		this._serialize(root.Left, sb)
		this._serialize(root.Right, sb)
	}

}

func (this Codec) deserialize(data string) *TreeNode {
	//获取字符串对应的切片
	nodes := make([]string, 0)
	for _, s := range strings.Split(data, this.SEP) {
		nodes = append(nodes, s)
	}
	return this._deserialize(&nodes)
}

func (this Codec) _deserialize(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	//1.获取前序遍历的第一个节点作为根节点
	first := (*nodes)[0]
	//2.更新接下来的左右子节点操作的切片
	*nodes = (*nodes)[1:]
	if first == this.NULL {
		return nil
	}

	val, _ := strconv.Atoi(first)
	root := &TreeNode{Val: val}

	root.Left = this._deserialize(nodes)
	root.Right = this._deserialize(nodes)
	return root
}
