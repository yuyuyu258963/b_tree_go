package main

//节点中的数据
type Item struct {
	key   int
	value interface{}
}

// B树
type BTree struct {
	m    int // number of leaf nodes
	root *BNode
}

// 树节点
type BNode struct {
	parent   *BNode // 指向父节点
	items    []*Item
	children []*BNode
}

// 初始化一个B树
// {m} max number of leaf children
func InitBtree(m int) *BTree {
	if m < 2 {
		panic("m must >= 1")
	}
	return &BTree{m: m}
}

// 获得初始化的节点
func InitBNode(parent *BNode, maxCap int) *BNode {
	return &BNode{
		parent:   parent,
		items:    make([]*Item, 0, maxCap),
		children: make([]*BNode, 0, maxCap), // 这里先预留一个位置方便后面节点的分裂
	}
}

// 拷贝数据
func (i *Item) Copy() *Item {
	if i == nil {
		return nil
	}
	return &Item{
		key:   i.key,
		value: i.value,
	}
}

// 判断是否为终端结点
func (N *BNode) IsLeaf() bool {
	return N.items == nil || len(N.children) == 0
}

// 融合节点
func (N *BNode) Merge(newNode *BNode) {
	N.items = append(N.items, newNode.items...)
	N.children = append(N.children, newNode.children...)
}

//不能向兄弟 借结点时 去找兄弟结点中size最小的
func (N *BNode) FindMinSizeNeighbor(index int) (item *BNode) {
	if index < 0 || index > len(N.children) {
		panic("index out of range")
	}
	if index == 0 || len(N.children[index-1].items) >= len(N.children[index+1].items) {
		item = N.children[index-1]
	} else if index == len(N.children)-1 {
		item = N.children[index+1]
	}
	return
}

// 借用兄弟结点的关键字
func (N *BNode) FindFitNeighbor(index int, maxCap int, replaceItem Item) (lendItem *Item) {
	if index < 0 || index >= len(N.children) {
		panic("index out of range")
	}
	lNode, rNode := N.children[index], N.children[index+1]
	if len(lNode.items) <= 0 && len(rNode.items) <= 0 {
		return
	} else {
		if len(lNode.items) >= len(rNode.items) && len(lNode.items) >= maxCap/2 {
			lendItem = lNode.items[len(lNode.items)-1]
			lNode.items = lNode.items[:len(lNode.items)-1]
		} else if len(rNode.items) >= maxCap/2 {
			lendItem = rNode.items[0]
			rNode.items = rNode.items[1:]
		}
	}

	return
}
