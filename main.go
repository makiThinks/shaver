package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(math.Pow(z, 2)-x) >= 0.1 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(4.0))
	fmt.Println(sqrt(16.0))
}
