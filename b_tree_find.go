package main

import "errors"

// 适用于查找对应 [key] 数据所在位置，或者是所在子树的位置
func (node *BNode) FindInsertIndex(key int) (index int, err error) {
	if len(node.items) < 1 {
		return -1, errors.New("empty node items")
	}
	for index = 0; index < len(node.items) && key > node.items[index].key; index++ {
	}
	return
}

// 查找或插入时 数据所在节点
func (N *BNode) FindNode(key int) (node *BNode, index int, err error) {
	if len(N.items) <= 0 {
		err = errors.New("node is empty")
		return
	}
	idx, err := N.FindInsertIndex(key)
	if err != nil {
	} else if (idx == len(N.items) && len(N.children) <= idx) || (idx == 0 && len(N.children) == 0) || (idx < len(N.items) && key == N.items[idx].key) {
		// 这里返回了三种情况:
		// 1. 找到相等的节点
		// 2. 无叶子节点时 查找的关键字小于节点中所有关键字
		// 3. 无叶子节点时 查找的关键字大于节点中所有关键字
		node = N
		index = idx
	} else if idx < len(N.children) {
		node, index, err = N.children[idx].FindNode(key)
	} else {
		err = errors.New("find data error")
	}
	return
}

// 从节点出发查找对应的数据
func (N *BNode) Find(key int) (item *Item) {
	TNode, index, err := N.FindNode(key)
	if err == nil && index < len(TNode.items) && key == TNode.items[index].key {
		item = TNode.items[index]
	}
	return item
}
