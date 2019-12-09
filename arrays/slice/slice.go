package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3} //slice

	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:len(s1)]
	fmt.Println(s2)

	s3 := a2[1:5]
	fmt.Println(s3)
}
