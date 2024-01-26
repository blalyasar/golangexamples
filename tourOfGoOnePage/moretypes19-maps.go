package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m3 = map[string]Vertex{
	"a1": {12, 14},
	"a2": {15, 18},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.6, -74.36,
	}
	fmt.Println(m)

	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			12, 14,
		},
		"Google": Vertex{
			34, 122,
		},
	}
	fmt.Println(m2)
	fmt.Println(m3)

}
