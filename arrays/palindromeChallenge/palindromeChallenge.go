package main

import (
	"fmt"
	"strconv"
)

func palindrome(num int) bool {

	numStr := strconv.Itoa(num)
	// fmt.Print(len(numStr))
	for i, j := 0, len(numStr)-1; i < len(numStr) && j >= 0; i++ {
		if numStr[i] != numStr[j] {
			return false
		}
		if i == len(numStr)/2 || j == len(numStr)/2 {
			return true
		}
		j--
	}
	return true
}

func main() {
	fmt.Println(palindrome(323))
}
