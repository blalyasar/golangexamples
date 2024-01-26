package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["answer"] = 42
	fmt.Println("value ", m["answer"])

	m["answer"] = 48
	fmt.Println("value ", m["answer"])

	delete(m, "answer")
	fmt.Println("value ", m["answer"])

	v, ok := m["answer"]
	fmt.Println("value ", v, "Present", ok)

}
