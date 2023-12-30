package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboligado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboligado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}

	carro1.ligarTurbo()

	// NOTAR USO DO & para utilizar uma variavel de tipo interface.
	var carro2 esportivo = &ferrari{"F40", false, 0}

	carro2.ligarTurbo()

	fmt.Println(carro2)
}
