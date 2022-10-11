package main

import "fmt"

func test() {
	t := InitBtree(3)
	t.Insert(&Item{1, "a"})
	t.Insert(&Item{3, "a"})
	t.Insert(&Item{0, "a"})
	t.Insert(&Item{5, "a"})
	t.Insert(&Item{6, "a"})
	t.Insert(&Item{7, "a"})
	t.Insert(&Item{8, "a"})
	t.Insert(&Item{9, "a"})
	t.Insert(&Item{10, "a"})
	t.Insert(&Item{-1, "a"})
	t.root.View()
	fmt.Println(t.Delete(7))
	fmt.Println(t.Delete(6))
	t.root.View()
	fmt.Println(t.Delete(5))
	t.root.View()
	t.Insert(&Item{4, "a"})
	t.root.View()
}

func main() {
	test()
}
