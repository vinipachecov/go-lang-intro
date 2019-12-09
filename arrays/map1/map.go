package main

import "fmt"

func main() {
	// var aprovados map[int]string  mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[123456789] = "Joao"
	aprovados[123456782] = "Maria"

	for cpf, nome := range aprovados {
		fmt.Println(nome, cpf)
	}

	delete(aprovados, 123456789)
	fmt.Println(aprovados)
}
