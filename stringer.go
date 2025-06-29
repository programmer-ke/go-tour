package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"John", 42}
	z := Person{"Zaphod", 9001}

	pa := &a
	pz := &z
	fmt.Println(a, z)
	fmt.Println(pa, pz)

	fmt.Printf("%T: %v, %T, %v\n", pa, pa, pz, pz)
}
