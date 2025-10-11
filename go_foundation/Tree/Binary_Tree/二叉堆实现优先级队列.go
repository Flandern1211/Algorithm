package main

import (
	"container/heap"
	"errors"
)

type SimpleMinPQ struct {
	//底层用数组实现
	head []int
	//堆中的元素
	size int
}

// 1.为什么要实现父，子节点的方法?
// 因为要先构成可以(得到父节点的)完全二叉树的结构才能实现二叉堆
// 父节点的索引
func (pq *SimpleMinPQ) parent(node int) int {
	return (node - 1) / 2
}

// 左子点的索引
func (pq *SimpleMinPQ) left(node int) int {
	return node*2 + 1
}

// 右子点的索引
func (pq *SimpleMinPQ) right(node int) int {
	return node*2 + 2
}

//2. 以下才是二叉堆的实现方法

// 交换元素的值,主要用于swim和sink函数
func (pq *SimpleMinPQ) swap(i, j int) {
	pq.head[i], pq.head[j] = pq.head[j], pq.head[i]
}

// 向堆中增加数据
func (pq *SimpleMinPQ) push(x int) {
	pq.head[pq.size] = x
	pq.swim(pq.size)
	pq.size++
}

// 向堆中删除数据
func (pq *SimpleMinPQ) pop() int {
	min := pq.head[0]
	pq.swap(0, pq.size-1)
	pq.size--
	pq.sink(0)
	return min
}

// 查，返回堆顶元素
func (pq *SimpleMinPQ) peek() int {
	return pq.head[0]
}

//上浮操作

func (pq *SimpleMinPQ) swim(node int) {
	for node > 0 && pq.head[pq.parent(node)] > pq.head[node] {
		pq.swap(pq.parent(node), node)
		node = pq.parent(node)
	}
}

// 下沉操作
func (pq *SimpleMinPQ) sink(node int) {
	for pq.left(node) < pq.size || pq.right(node) < pq.size {
		//min := pq.head[node]
		//if pq.head[pq.left(node)] < min{
		//	min = pq.head[pq.left(node)]
		//}
		//if pq.head[pq.right(node)] < min{
		//	min = pq.head[pq.right(node)]
		//}
		min1 := node
		if pq.left(node) < pq.size && pq.head[pq.left(node)] < pq.head[min1] {
			min1 = pq.left(node)
		}
		if pq.right(node) < pq.size && pq.head[pq.right(node)] < pq.head[min1] {
			min1 = pq.right(node)
		}

		if min1 == node {
			break
		}
		pq.swap(node, min1)
		node = min1
	}
}

type MyPriorityQueue struct {
	//堆数组
	head []interface{}
	//堆中元素的数量
	size int

	//元素比较器
	comparator func(x, y interface{}) int
}

// 构造函数
func NewMyPriorityQueue(capacity int, comparator func(x, y interface{}) int) *MyPriorityQueue {
	return &MyPriorityQueue{
		head:       make([]interface{}, capacity),
		size:       0,
		comparator: comparator,
	}
}

// 返回堆的大小
func (pq *MyPriorityQueue) Size() int {
	return pq.size
}

// 判断堆是否为空
func (pq *MyPriorityQueue) IsEmpty() bool {
	return pq.size == 0
}

// 父节点的索引
func (pq *MyPriorityQueue) Parent(node int) int {
	return (node - 1) / 2
}

func (pq *MyPriorityQueue) Left(node int) int {
	return node*2 + 1
}

func (pq *MyPriorityQueue) Right(node int) int {
	return node*2 + 2
}

// 交换数组的两个元素
func (pq *MyPriorityQueue) Swap(i, j int) {
	pq.head[i], pq.head[j] = pq.head[j], pq.head[i]
}

func (pq *MyPriorityQueue) Peek() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	return pq.head[0], nil
}

func (pq *MyPriorityQueue) Push(x interface{}) {
	//扩容
	if pq.size == len(pq.head) {
		pq.resize(2 * len(pq.head))
	}
	//新的元素追加到最后
	pq.head[pq.size] = x
	pq.swim(pq.size)
	pq.size++
}

// 删， 删除堆顶元素，
func (pq *MyPriorityQueue) Pop() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	res := pq.head[0]
	//把堆底元素放在堆顶
	pq.Swap(0, pq.size-1)
	//避免对象游离
	pq.head[pq.size-1] = nil
	pq.size--
	//让后下降到正确位置
	pq.sink(0)
	if pq.size > 0 && pq.size == len(pq.head)/4 {
		pq.resize(len(pq.head) / 2)
	}
	return res, nil
}

// 上浮操作
func (pq *MyPriorityQueue) swim(node int) {
	for node > 0 && pq.comparator(pq.head[pq.Parent(node)], pq.head[node]) > 0 {
		pq.Swap(pq.Parent(node), node)
		node = pq.Parent(node)
	}
}

// 下沉操作
func (pq *MyPriorityQueue) sink(node int) {
	for pq.Left(node) < pq.size {
		//比较自己和左右节点的大小
		minNode := node
		if pq.Left(node) < pq.size && pq.comparator(pq.head[pq.Left(node)], pq.head[minNode]) > 0 {
			minNode = pq.Left(node)
		}
		if pq.Right(node) < pq.size && pq.comparator(pq.head[pq.Right(node)], pq.head[minNode]) > 0 {
			minNode = pq.Right(node)
		}
		if minNode == node {
			break
		}

		pq.Swap(node, minNode)
		//上一层的比较后比较下一层的
		node = minNode
	}

}
