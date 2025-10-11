package HashArray

import "math/rand"

// 哈希数组的实现
type MyArrayHashMap struct {
	//存储key和其在arr中的索引
	m map[int]int
	//在数组中存储对应的Node key和value
	//为什么还要存储key值呢？
	arr []Node
}

// 为什么要用Node来存储呢？就一个value不就行了？
// 因为要实现随机取key值，需要在数组中随机而不是在map中，所以数组中也必须存一个key
type Node struct {
	key   int
	value int
}

// 创建和初始化一个哈希数组
// 为什么这个就返回*，哈希链表返回的不是*呢？
func NewMyArrayHashMap() *MyArrayHashMap {
	return &MyArrayHashMap{
		m:   make(map[int]int),
		arr: make([]Node, 0),
	}
}

// 实现Get方法
func (this *MyArrayHashMap) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	//通过this.m获取其索引，然后根据索引来得到Node，再获取其中的value值
	return this.arr[this.m[key]].value
}

// 实现Put方法，增加/修改
func (this *MyArrayHashMap) Put(key int, val int) {
	//1.如果存在的话，就直接修改key对应索引的值
	if _, ok := this.m[key]; ok {
		this.arr[this.m[key]].value = val
		return
	}
	//2.该key不存在，则直接添加，首先在数组中添加，也就是直接添加到最后一位，map对应key的索引
	//也就是最后一位的索引了
	node := Node{key: key, value: val}
	this.arr = append(this.arr, node)
	this.m[key] = len(this.arr) - 1
}

//实现Remove方法，删除指定的key及其对应的valuer
//（将需删节点与数组最后一位交换删除，可以直接得到两个节点，一个是map存的索引得到，一个根据len-1得到）O（1）
//删除数组最后一位节点也为o（1）

func (this *MyArrayHashMap) Remove(key int) bool {
	//1.如若不存在，则直接返回true当作删除成功即可
	if _, ok := this.m[key]; !ok {
		return true
	}
	//2.该key存在的情况
	//先得到其对应索引
	index := this.m[key]
	//根据索引找到对应节点
	node := this.arr[index]
	//获取数组最后一位的节点
	e := this.arr[len(this.arr)-1]
	//两个节点交换
	this.arr[index] = e
	this.arr[len(this.arr)-1] = node
	//更新该节点在map中的索引
	this.m[e.key] = index
	//删除指定节点在数组中的节点，和map中对应的索引
	this.arr = this.arr[:len(this.arr)-1]
	delete(this.m, node.key)
	return true

}

// 实现随意获取一个key值的方法
func (this *MyArrayHashMap) randomKey() int {
	n := len(this.arr)
	randomIndex := rand.Intn(n)
	return this.arr[randomIndex].key
}

// 判断哈希数组中是否存在该key
func (this *MyArrayHashMap) containKey(key int) bool {
	_, ok := this.m[key]
	return ok
}

// 返回哈希数组的长度
func (this *MyArrayHashMap) size() int {
	return len(this.m)
}
