package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1.precoComDesconto())
}
