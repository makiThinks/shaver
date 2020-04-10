// bootstrap of this project
package main

import (
	"fmt"
	"math"

	"maki.io/demo/shaver/util"
)

func sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(math.Pow(z, 2)-x) >= 0.1 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	s := util.ReverseString("go is weird")
	i := sqrt(16)
	fmt.Println(s)
	fmt.Println(i)
}
