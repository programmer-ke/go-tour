package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
	fmt.Printf("type i: %T\n", i)
	t := T{"hello"}
	t.M()
	fmt.Printf("type t: %T\n", t)
}
