package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	println("Iniciou rotina")
	time.Sleep(time.Second)

	c <- 1

	println("só depois que foi lido")
}

func main() {
	c := make(chan int) //channel sem buffer, bloqueia logo que um valor é recebido

	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido")
	// fmt.Println(<-c)// dead lock

}
