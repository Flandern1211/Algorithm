package main

import "container/list"

//hashmap的实现基石,存储key和value

type KVNode struct {
	key   int
	value int
}

// HashMap的底层实现数组
type ExampleChainingHashMap struct {
	//一个存储链表地址的指针数组
	table []*list.List
}

// 工厂函数创建新hashmap,并指定其容量
func NewChainingHashMap(cap int) ExampleChainingHashMap {
	return ExampleChainingHashMap{
		table: make([]*list.List, cap),
	}
}

// 哈希函数，用于使key始终可以存入数组中
func (h *ExampleChainingHashMap) hash(key int) int {
	return key % len(h.table)
}

// get查
func (h *ExampleChainingHashMap) Get(key int) (int, bool) {
	index := h.hash(key)
	if h.table[index] == nil {
		return -1, false
	}

	for e := h.table[index].Front(); e != nil; e.Next() {
		node := e.Value.(KVNode)
		if node.key == key {
			return node.value, true
		}

	}

	//链表中没有目标key
	return -1, false
}

// 增
func (h *ExampleChainingHashMap) Add(key int, value int) bool {
	index := h.hash(key)
	if h.table[index] == nil {
		h.table[index] = list.New()
		h.table[index].PushBack(KVNode{key, value})
		return true
	}

	for e := h.table[index].Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode)
		if node.key == key {
			node.value = value
			return true
		}
	}
	//链表中无目标key
	//这里使用addFirst或addLast都可以
	h.table[index].PushBack(KVNode{key, value})
}

//删

func (h *ExampleChainingHashMap) Del(key int) (int, bool) {
	index := h.hash(key)
	if h.table[index] == nil {
		return -1, false
	}

	for e := h.table[index].Front(); e != nil; e = e.Next() {
		node := e.Value.(KVNode)
		if node.key == key {
			value := node.value
			h.table[index].Remove(e)
			return value, true
		}
	}

}
