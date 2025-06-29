package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.2f", e)
}

func abs(x float64) float64 {
	if x < 0 {
		return x * -1
	} else {
		return x
	}

}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	for {
		prev := z
		z -= (z*z - x) / (2 * z)
		diff := z - prev
		if abs(diff) < 1e-10 {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	_, err := Sqrt(-2)
	fmt.Printf("%T\n", err)
}
