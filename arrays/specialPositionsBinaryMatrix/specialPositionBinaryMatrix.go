package main

import "fmt"

func numSpecial(mat [][]int) int {
	specialValues := 0
	foundSpecialValue := false
	// main loop
	for i, _ := range mat {
		for j, _ := range mat[i] {
			fmt.Println("value: ", i, j)
			// // spot special value
			if mat[i][j] == 1 && foundSpecialValue != true {
				foundSpecialValue = true
				fmt.Println("Found special value: ", i, j)
				fmt.Println("Column check")

				// check columns
				for k, _ := range mat {
					fmt.Println("mat[k][j]", mat[k][j])
					if k != i && mat[k][j] == 1 {
						foundSpecialValue = false
						break
					}
				}
				// check row
				fmt.Println("Row check")
				for l, _ := range mat[i] {
					if l != j {
						if mat[i][l] == 1 {
							foundSpecialValue = false
							break
						}
					}
				}

				if foundSpecialValue == true {
					specialValues = specialValues + 1
					foundSpecialValue = false
				}
			}
		}
	}
	return specialValues
}
func main() {
	// res := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	res := [][]int{{0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}
	// len(res[0])
	fmt.Println("Special values = ", numSpecial(res))
}
