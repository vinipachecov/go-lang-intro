package main

import "fmt"
import "math"

func main() {
	const PI float64 = 3.1415
	var raio = 3.2
	area := PI * math.Pow(raio, 2)
	fmt.Println("Area = ", area);
	const (
		a = 2
		b = 1
	)
	fmt.Println(a,b)
	g,h,i := 2, 3, "epa"
	fmt.Print(g,h,i)
}