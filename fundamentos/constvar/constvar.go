package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)
	fmt.Println("Area = ", area)
	const (
		a = 2
		b = 1
	)
	area = 50
	fmt.Println("Area = ", area)
	fmt.Println(a, b)
	g, h, i := 2, 3, "epa"
	fmt.Print(g, h, i)
}
