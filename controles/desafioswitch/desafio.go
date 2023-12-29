package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n <= 7:
		return "b"
	case n >= 6 && n <= 5:
		return "C"
	case n >= 4 && n <= 3:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(5.8))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(1.8))
}
