package main

// 查找数值
func (t *BTree) Find(key int) (item *Item) {
	item = t.root.Find(key)
	return item.Copy()
}

// 向树中插入节点
func (t *BTree) Insert(item *Item) {
	if t.root == nil {
		root := InitBNode(nil, t.m)
		root.items = append(root.items, item.Copy())

		t.root = root
	} else {
		BTNodeInsert(t, item)
	}
}

// 删除B树中的结点
func (t *BTree) Delete(key int) (item *Item) {
	item = DeleteItem(t, key)
	return
}
