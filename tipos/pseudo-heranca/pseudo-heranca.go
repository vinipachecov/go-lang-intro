package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro // campos anonimos
}

func main() {
	f := ferrari{}
	f.nome = "Nome"
	f.velocidadeAtual = 200

	fmt.Println(f)

}
