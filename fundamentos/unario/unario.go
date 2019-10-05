package main

import "fmt"

func main() {
	x := 1
	y := 2

	// posfixação

	x++
	fmt.Println(x)
	y--
	fmt.Println(y)

	// fmt.Println(x == y--) não funciona em go
}
