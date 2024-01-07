package main

func main() {
	ch := make(chan int, 1)

	// assinging value
	ch <- 10

	// consuming value from channel
	println(<-ch)

}
