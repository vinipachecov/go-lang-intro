package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func printNum(num int) {
	fmt.Println(num)
}

func main() {
	p1, p2 := 1, 2
	primeiro, segundo := trocar(p1, p2)

	println(primeiro, segundo)

	for i := 0; i < 1000000; i++ {
		go printNum(i)
	}
}
