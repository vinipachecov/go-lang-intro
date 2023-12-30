package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:id`
	Nome  string   `json:nome`
	Preco float64  `json:preco`
	Tags  []string `json:tags`
}

func main() {
	p1 := produto{1, "Notebook", 1899.90, []string{"Promocao", "Eletronico"}}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{ "id": 2, "nome": "Caneta", "tags": ["Papelaria", "Importado"]}`

	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags)
}
