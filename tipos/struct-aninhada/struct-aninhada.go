package main

import "fmt"

type Item struct {
	productid  int
	quantidade int
	preco      float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := Pedido{
		userID: 1,
		itens: []Item{
			Item{1, 2, 12.10},
			Item{1, 20000, 40.10},
			Item{1, 2, 10.10},
		},
	}

	fmt.Println(pedido.valorTotal())

}
