package main

import "time"

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second / 2)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	channel := make(chan int)

	go doisTresQuatroVezes(2, channel)

	a, b := <-channel, <-channel

	println(a, b)

	println(<-channel)

	// listen to a value that won't be sent will result in a deadlock
	// println(<-channel)
}
