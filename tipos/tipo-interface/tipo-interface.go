package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}

	fmt.Println(coisa)

	coisa = 3

	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"

	fmt.Println(coisa2)
}
