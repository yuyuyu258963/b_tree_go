package main

import (
	"errors"
)

// 找到比key大的数中的最小数的位置
func FindIndex(arr []*Item, key int) (index int) {
	for index = 0; len(arr) > index && arr[index].key < key; index++ {
	}
	return
}

// 处理上溢出 分裂
func (N *BNode) SplitNode(t *BTree) {
	var p *BNode
	maxCap := t.m
	currentNode := N
	for len(currentNode.items) >= maxCap {
		insertIndex := 0
		p = currentNode.parent
		mid := len(N.items) / 2
		// 若当前节点为根节点时 需创建新的根节点
		if p == nil {
			newRoot := InitBNode(nil, t.m)
			t.root, p = newRoot, newRoot
		} else {
			// 找到在父节点插入的位置
			insertIndex = FindIndex(p.items, currentNode.items[mid].key)
		}
		p.items = append(p.items[:insertIndex], append([]*Item{currentNode.items[mid]}, p.items[insertIndex:]...)...)
		// 更新父节点的子树
		lNode := &BNode{
			parent:   p,
			items:    currentNode.items[:mid],
			children: make([]*BNode, 0, maxCap),
		}
		rNode := &BNode{
			parent:   p,
			items:    currentNode.items[mid+1:],
			children: make([]*BNode, 0, maxCap),
		}
		if mid < len(currentNode.children) {
			// 更新分裂后节点对应父节点的指向
			for _, child := range currentNode.children[:mid+1] {
				child.parent = lNode
			}
			for _, child := range currentNode.children[mid+1:] {
				child.parent = rNode
			}
			lNode.children = currentNode.children[:mid+1]
			rNode.children = currentNode.children[mid+1:]
		}
		// 更新插入后的叶子节点
		if insertIndex < len(p.children) {
			p.children = append(p.children[:insertIndex],
				append([]*BNode{lNode, rNode}, p.children[insertIndex+1:]...)...)
		} else {
			p.children = append(p.children[:insertIndex], []*BNode{lNode, rNode}...)
		}

		currentNode = p
	}
}

// 处理向B-Tree中插入数据
func BTNodeInsert(t *BTree, item *Item) (err error) {
	N := t.root
	// 定位
	TNode, index, err := N.FindNode(item.key)
	if err != nil {
		return
	} else if index < len(TNode.items) && TNode.items[index].key == item.key {
		return errors.New("key can't be repeat")
	}
	// 插入
	if index < len(TNode.items) {
		TNode.items = append(TNode.items[:index], append([]*Item{item}, TNode.items[index:]...)...)
	} else {
		TNode.items = append(TNode.items, item)
	}

	TNode.SplitNode(t)
	return
}
