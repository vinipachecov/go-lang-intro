package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		// time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "vendo pao", 3)
	// fale("João", "vendo 2 pao", 1)

	go fale("Maria", "vendo pao", 500)
	go fale("João", "vendo 2 pao", 500)
	time.Sleep(time.Second * 10)
}
