package main
import "fmt"

type Number interface{
	int64 | float64
}

func SumInts(m map[string]int64) int64{
	var s int64
	for _, v := range m{
		s += v
	}
	return s
}

func SumFloats(m map[string]float64 ) float64{
	var s float64
	for _,v :=  range m {
		s += v
	}
	return s
}

func SumIntOrFloats[K comparable, V int64   | float64 ](m map[K] V) V{
	var s V
	for _, v := range m{
		s += v
	}
	return s
}


func main() {
	ints := map[string]int64{
		"first": 35,
		"second": 26,
	}

	floats := map[string]float64{
		"first": 35.2,
		"second": 26.2,
	}
	fmt.Printf("Non Generic SUms +v %v\n",SumInts(ints), SumFloats(floats))

	fmt.Printf("\nGeneric SUMms +v %v\n", 
	SumIntOrFloats[string, int64 ](ints),
	SumIntOrFloats[string, float64 ](floats))


}