package main

import "fmt"

// 查看B树
func (N *BNode) View() {
	currentNode := N
	var g = []*BNode{currentNode}
	var preLen int
	fmt.Printf("\n================================================================\n")
	for {
		preLen = len(g)
		for _, d := range g {
			if d != nil {
				fmt.Printf("|")
				showItems(d)
				fmt.Printf("|")

				if len(d.children) > 0 {
					g = append(g, d.children...)
				}
			}
		}
		g = g[preLen:]
		fmt.Println()
		if len(g) == 0 {
			break
		}
	}
	fmt.Printf("================================================================\n")
}

func showItems(N *BNode) {
	for _, d := range N.items {
		// fmt.Printf(" point:%v d: %+v parent:%v  ", &N, *d, &N.parent)
		fmt.Printf("%+v", *d)
	}
}
