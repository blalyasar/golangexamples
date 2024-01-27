//Choosing a value or pointer receiver

package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v *vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &vertex{3, 4}

	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

}
