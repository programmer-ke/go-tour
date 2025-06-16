package main

import "fmt"

func abs(x float64) float64 {
	if x < 0 {
		return x * -1
	} else {
		return x
	}

}

func Sqrt(x float64) float64 {
	// finds square root using newton's method
	z := x / 2
	for {
		prev := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		diff := z - prev
		if abs(diff) < 1e-10 {
			fmt.Println(diff)
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(9))
}
