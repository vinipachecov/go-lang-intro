package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("O byte é do tipo", reflect.TypeOf(b))
	x := 12.30
	fmt.Print("Tipo do x =", reflect.TypeOf(x))
}
