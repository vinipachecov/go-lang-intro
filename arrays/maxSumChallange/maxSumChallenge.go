package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	// var max1 = 0
	// var max2 = 0

	sort.Ints(nums)

	// find two highest values considering num - index
	// for _, number := range nums {
	// 	if max1 == 0 && max2 == 0 {
	// 		max1 = number - 1
	// 	}

	// 	if max1 != 0 && max1 < number-1 {
	// 		max1 = number - 1
	// 	}

	// 	if max2 == 0 && max2 < number-1 {
	// 		max2 = number - 1
	// 	} else if max2 < number-1 && max1 > number-1 {
	// 		max2 = number - 1
	// 	}
	// }

	// fmt.Println("max1 ", max1)
	// fmt.Println("max1 ", max2)

	fmt.Print(nums)
	return nums[0] * nums[1]
}

func main() {
	result := maxProduct([]int{3, 4, 5, 2})
	fmt.Print("Result: ", result)
}
