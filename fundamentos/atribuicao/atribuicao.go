package main

import "fmt"

func main() {
	// declaração
	var b byte = 3
	fmt.Println(b)

	// inferencia de tipo
	i := 3
	i += 3

	i -= 3
	i /= 2
	i *= 2
	i %= 2

	x, y := 1, 2

	// swap
	x, y = y, x
	fmt.Println(b)
}
