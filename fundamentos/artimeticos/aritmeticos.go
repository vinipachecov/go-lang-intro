package main

import (
	"fmt"
	"math"
)
func main() {
	a :=3
	b :=2

	fmt.Println("Soma =>", a+b)	
	fmt.Println("Subtração =>", a+b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a+b)
	fmt.Println("Modulo =>", a%b)

	fmt.Println("E =>", a&b);
	fmt.Println("OU =>", a|b);
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("MAIOR =>", math.Max(float64(a), float64(b)))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
