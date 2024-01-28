package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Art ", 42}
	z := Person{"Zaa ", 23}
	fmt.Println(a, z)
}
