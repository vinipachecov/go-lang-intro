package main

import "fmt"

func main() {
	i := 1

	// go não tem aritmética de ponteiros
	var p *int = nil
	// p agora tem o endereço do valor i
	p = &i
	fmt.Println(p)

	*p++
	i++

	fmt.Println(p, *p, i, &i)

}
