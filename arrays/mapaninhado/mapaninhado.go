package main

import "fmt"

func main() {
	funsporLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1546.75,
			"Guga Pereira":   13000,
		},
		"j": {
			"Juga Pereira":   13000,
			"Jabriela Silva": 1546.75,
		},
		"V": {
			"Vinicius": 5000,
		},
	}

	for letra, funs := range funsporLetra {
		fmt.Println(letra, funs)
	}
}
