package main

import "fmt"

func somar(a float64, b float64) float64 {
	return a + b
}

func imprimir(valor float64) {
	fmt.Println(valor)
}

func main() {
	imprimir(somar(2,2))
}