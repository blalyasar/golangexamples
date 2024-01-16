package main
// go get golang.org/x/example/hello/reverse

import (
	"fmt"
	"golang.org/x/example/hello/reverse"

)

func main() {
	// fmt.Println(reverse.String("Hello"))
	fmt.Println(reverse.String("Hello"), reverse.Int(24601) )
}

// /004-workspace$ go work init ./hello/
