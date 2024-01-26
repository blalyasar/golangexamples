package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var pow2 = []int{11, 12, 13, 14, 15, 16, 17, 18}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow3 := make([]int, 10)
	for i := range pow3 {
		pow3[i] = 1 << uint(i) // 2**i
	}
	for _, value := range pow3 {
		fmt.Printf("%d\n", value)
	}

}
