package main

import "fmt"

func Reverse(s string) string {
	b := []byte(s)
	for i,j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1{
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}



func main(){
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
//	doubleRev := Reverse(rev)
	fmt.Printf("hola hola \n" )
	fmt.Printf("orig %q\n", input)
	fmt.Printf("rev %q\n",  rev)
	// fmt.Printf("rev again %q\n ",  doubleRev)
}