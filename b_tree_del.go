package main

// 处理节点的下溢
// 传入结点的最小限制
func (N *BNode) SolveUnderFlow(t *BTree, maxCap int) {
	var placeInFatherNode = 0
	p := N.parent
	// 当前结点为根节点 且无数据
	if p == nil || len(N.items) >= maxCap/2 {
		if len(N.items) == 0 {
			t.root = N.children[0]
		}
		return
	}

	// 找到节点在兄弟中的位置
	for index, child := range p.children {
		if child == N {
			placeInFatherNode = index
			break
		}
	}

	// 旋转
	if placeInFatherNode > 0 {
		// 向左兄弟借
		lNode := p.children[placeInFatherNode]
		if len(lNode.items) > maxCap/2 {
			templateItem := lNode.items[len(lNode.items)-1]
			// 进行右旋
			N.items = append([]*Item{p.items[placeInFatherNode]}, N.items...)
			p.items[placeInFatherNode] = templateItem
			lNode.items = lNode.items[:len(lNode.items)-1]
			// 修改旋转后结点的父结点的指向
			lNode.children[len(lNode.children)-1].parent = N
			N.children = append([]*BNode{lNode.children[len(lNode.children)-1]}, N.children...)
			lNode.children = lNode.children[:len(lNode.children)-1]
		}
	}
	if placeInFatherNode < len(p.items) {
		// 向右兄弟借
		rNode := p.children[placeInFatherNode+1]
		if len(rNode.items) > maxCap/2 {
			templateNode := rNode.items[0]
			// 进行左旋
			N.items = append(N.items, templateNode)
			p.items[placeInFatherNode] = templateNode
			rNode.items = rNode.items[1:]

			rNode.children[0].parent = N
			N.children = append(N.children, rNode.children[0])
			rNode.children = rNode.children[1:]
		}
	}
	//合并
	if placeInFatherNode > 0 {
		merge(p, p.children[placeInFatherNode-1], N, placeInFatherNode-1)
	} else {
		merge(p, N, p.children[placeInFatherNode+1], placeInFatherNode)
	}

	p.SolveUnderFlow(t, maxCap)

}

// 跟前驱交换数据
// ``return: 最后抽取结点``
func (N *BNode) predecessor(pos int) (pred *BNode) {
	pred = N.children[pos]
	for !pred.IsLeaf() {
		pred = pred.children[len(pred.children)-1]
	}
	swapIndex := len(pred.items) - 1
	// 前驱结点删除已经交换的数据项
	N.items[pos] = pred.items[swapIndex]
	pred.items = pred.items[:swapIndex]
	return pred
}

// 合并兄弟结点
func merge(p, l, r *BNode, pos int) {
	l.items = append(l.items, p.items[pos])
	p.items = append(p.items[:pos], p.items[pos+1:]...)
	p.children = append(p.children[:pos+1], p.children[pos+2:]...)
	l.Merge(r)
}

// 从树种删除元素
func DeleteItem(t *BTree, key int) (item *Item) {
	TNode, index, err := t.root.FindNode(key)
	// 找到该结点就继续， 否则退出
	if err == nil && index < len(TNode.items) && key == TNode.items[index].key {
		item = TNode.items[index]
	} else {
		return
	}
	isLeaf := TNode.IsLeaf()
	// 为最低层叶子节点的时候直接删除
	if isLeaf {
		TNode.items = append(TNode.items[:index], TNode.items[index+1:]...)
	} else {
		// 若不是叶子结点则不断地从前驱结点抽取数据
		TNode = TNode.predecessor(index)
	}
	//
	TNode.SolveUnderFlow(t, t.m)
	return
}
