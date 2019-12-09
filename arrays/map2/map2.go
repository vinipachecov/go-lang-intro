package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"jose":     11000,
		"gabriela": 5000,
		"mauro":    620,
	}

	funcsESalarios["Rafael"] = 1350
	fmt.Println("test")
	delete(funcsESalarios, "iNexistente")
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
