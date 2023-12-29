package main

import "fmt"

func obterValorAprovado(numero int) int {
	// Line 7 with "defer" will execute when this function returns or finishes
	defer fmt.Println("Fim!")
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return 500
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
}
