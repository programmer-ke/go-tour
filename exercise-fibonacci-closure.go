package main

import "fmt"

func fibonacci() func() int {

	var curr, next = 0, 1

	return func() int {
		result := curr
		curr, next = next, curr+next
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
