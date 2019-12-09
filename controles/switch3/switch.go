package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "INTEIRO"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei."

	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo("here"))
	fmt.Println(tipo(1))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo("here"))

}
