package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[T]) String() string {
	if l.next == nil {
		return fmt.Sprintf("%v", l.val)
	}
	return fmt.Sprintf("%v -> %s", l.val, l.next)
}

func main() {
	node := List[int]{}
	node.val = 1
	node.next = &List[int]{nil, 2}
	node.next.next = &List[int]{val: 3}
	fmt.Println(node)
	fmt.Println(*node.next)
}
