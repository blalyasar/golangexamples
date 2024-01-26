package main

import (
	"fmt"
	"math"
)
// https://go.dev/tour/methods/7

type vertex struct {
	x, y float64
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func AbsFunc(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &vertex{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
