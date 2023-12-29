package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	val := 10

	inc1(val)

	fmt.Println(val)

	inc2(&val)

	fmt.Println(val)
}
