package HashLink

// 实现哈希链表，既可以在遍历hashmap时，按照存入的顺序来遍历出来

// 这是哈希链表，是我们自己来定义实现的，所以使用创建结构体的方式来创建
type MyLinkedHashMap struct {
	//为什么要使用双链表来实现呢？ ——主要是为了好遍历和好删除，使得不会增加遍历的时间复杂度

	head *Node            //指向链表的头结点
	tail *Node            //指向链表的尾节点
	m    map[string]*Node //对应的key直接指向链表的节点地址，这样就可以直接寻到值，不增加复杂度
}

// 实现链表中的单个节点，链表就是一个一个节点相连的
type Node struct {
	key string //为什么还要存个key呢？ ——如果只有map中有key的话，想遍历的时候还是无序的，但链表是绝对有序的
	//所以必须存key才能实现按存入顺序来遍历

	val  int
	next *Node
	prev *Node
}

//为什么不用再专门实现一个链表呢？
//1.在这个实现中链表只是作为哈希链表的一个组成部分
//2.而且通过对head和tail虚拟节点的操作和方法就足够我们利用它的特性实现遍历了

// 便于用户2直接创建和初始化一个哈希链表
func Constructor() MyLinkedHashMap {
	head := &Node{}         //创建一个空节点，将其地址赋值给head，这个节点将永远作为链表的头结点（虚拟头节点，便于链表操作）
	tail := &Node{}         //创建一个新的Node节点，并获取它的地址赋值给tail变量，这个节点将作为链表的尾节点（虚拟尾节点，便于链表操作）
	head.next = tail        // 将链表头尾节点链接，成为一个链表
	tail.prev = head        // 将尾节点的前驱指向头节点
	return MyLinkedHashMap{ //返回初始化完的哈希链表
		head: head,
		tail: tail,
		m:    make(map[string]*Node), //直接创建一个map作为底层map
	}
}

// 实现哈希链表的Get方法，根据key得到value
func (this *MyLinkedHashMap) Get(key string) int {
	//this.m[key]直接使用底层map来找到该节点，存在ok为true，
	//直接返回该节点的val字段的值，node.val是系统自动给你解引用了
	if node, ok := this.m[key]; ok {
		return node.val
	}
	//返回-1表示不存在
	return -1
}

// 实现哈希链表的存储方法,存储key和对应的val值，也可修改key对应的val值
func (this *MyLinkedHashMap) Put(key string, val int) {
	//先判断该key是否存在，若不存在，则需要创建一个新的节点存储key和val值
	//并将该节点通过addLastNode方法加到虚拟的链表尾部
	if _, ok := this.m[key]; !ok {
		node := &Node{key: key, val: val}
		this.addLastNode(node)
		this.m[key] = node
		return
	}
	//若存在的话直接修改该节点的val值即可
	this.m[key].val = val
}

// 实现哈希链表的Remove方法，可以通过给定key值来删除哈希链表中对应的key和val值
func (this *MyLinkedHashMap) Remove(key string) {
	//首先判断该key值是否存在，不存在直接返回即可
	if _, ok := this.m[key]; !ok {
		return
	}
	//底层map的key和虚拟链表中的node节点都要删除
	//先获取该key对应的node节点
	node := this.m[key]
	//删除map中的对应的key
	delete(this.m, key)
	//删除虚拟链表中的该节点
	this.removeNode(node)
}

// 判断哈希链表中是否存在该key
func (this *MyLinkedHashMap) ContainsKey(key string) bool {
	_, ok := this.m[key]
	return ok
}

// 得到哈希链表中的所有key，且是按照存入顺序来遍历出来的（虚拟链表最大的作用就在这里）
func (this *MyLinkedHashMap) Keys() []string {
	//创建一个类型为string的切片
	keyList := make([]string, 0)
	//从头结点遍历到尾节点
	for p := this.head.next; p != this.tail; p = p.next {
		keyList = append(keyList, p.key)

	}
	//为什么还专门建个切片来？不是直接遍历链表输出就行了吗?
	//返回统一格式：返回切片使得API更一致，调用者可以得到一个标准的数组格式数据
	//便于后续处理：切片可以方便地进行排序、过滤等操作
	//更好的封装性：隐藏了内部链表结构的实现细节
	//灵活性：调用者可以根据需要决定如何使用这个键列表（遍历、索引访问等）
	return keyList
}

// 对链表的操作
// 1.向链表尾部加节点
func (this *MyLinkedHashMap) addLastNode(x *Node) {
	//由尾部节点得到其上一节点
	temp := this.tail.prev
	//将x节点连接到尾节点（虚拟尾节点）
	x.next = this.tail
	//将x节点的连接到temp节点，也就是原先的虚拟尾节点的前节点
	x.prev = temp
	//temp节点连接x节点
	temp.next = x
	//更新虚拟尾节点的前节点并连接
	this.tail.prev = x
}

// 删除指定节点
func (this *MyLinkedHashMap) removeNode(x *Node) {
	//得到x节点的前和后节点
	prev := x.prev
	next := x.next
	//将这两个节点连起来
	prev.next = next
	next.prev = prev
	//将x节点的连接断掉，系统会自动清除该节点
	x.next = nil
	x.prev = nil

}
