package main

import "fmt"

// https://go.dev/tour/methods/6

type vertex struct {
	x, y float64
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func scalefunc(v *vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := vertex{3, 4}
	v.scale(2)
	scalefunc(&v, 10)

	p := &vertex{4, 3}
	p.scale(3)
	scalefunc(p, 8)
	fmt.Println(v, p)
}
