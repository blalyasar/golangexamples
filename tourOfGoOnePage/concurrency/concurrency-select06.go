package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(10 * time.Millisecond)
	boom := time.After(20 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("TICK")
		case <-boom:
			fmt.Println("BOOM")
			return
		default:
			fmt.Println("..DEF.")
			fmt.Println(50 * time.Millisecond)

		}
	}
}
