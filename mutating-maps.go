package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	m["Question"] = 0
	fmt.Println(m)

	m["Answer"] = 48
	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println(m)
	fmt.Println(m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:",v, "Present?", ok)

	value := m["Answer"]
	fmt.Println("missing value", value)
	
}
