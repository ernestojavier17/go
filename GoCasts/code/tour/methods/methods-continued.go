package main

import (
	"fmt"
	"math"
)

type MyFloat float64

//Method declare on a non-struct type
//You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}