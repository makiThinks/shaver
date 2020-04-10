// bootstrap of this project
package main

import (
	"fmt"
	"math"

	"maki.io/demo/shaver/service"
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
	// test part
	s := util.ReverseString("go is weird")
	i := sqrt(16)
	fmt.Println(s)
	fmt.Println(i)

	// the server part
	name := "testPage"
	p1 := &service.Page{Title: name, Body: []byte("This is a sample page")}
	save_err := p1.Save()
	if save_err != nil {
		fmt.Printf("step 1: page save error %v", save_err)
	} else {
		p2, load_err := service.LoadPage(name)
		if load_err != nil {
			fmt.Printf("step 2: page load error %v", load_err)
		} else {
			fmt.Println(string(p2.Body))
		}
	}

}
