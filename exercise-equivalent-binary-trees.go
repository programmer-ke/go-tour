package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var v1, v2 int
	var ok1, ok2 bool
	for {
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
		if !ok1 || !ok2 {
			break
		}
		if v1 != v2 {
			return false
		}
	}
	return ok1 == ok2
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
