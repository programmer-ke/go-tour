package main

import "fmt"

const Pi = 3.15

func main() {
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
