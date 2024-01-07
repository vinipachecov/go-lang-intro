package main

import "fmt"

func rotina(ch chan int) {
	fmt.Println("Executou")

	ch <- 1
	fmt.Println("Executou 1")
	ch <- 2
	fmt.Println("Executou 2")
	ch <- 3
	fmt.Println("Executou 3")
	ch <- 4
	fmt.Println("Executou 4")
	ch <- 5
	fmt.Println("Executou 5")
	ch <- 6
	ch <- 7
}

func main() {
	ch := make(chan int, 3)

	go rotina(ch)

	fmt.Println(<-ch)
}
