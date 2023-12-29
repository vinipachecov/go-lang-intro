package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID    string
	name  string
	email string
}

func main() {

	// const and var
	const dbURI = "http://www.google.com"

	// type inference :=
	str := "123456"

	fmt.Println(str)
	fmt.Println(dbURI)

	arr := [4]int{1, 2, 3, 4}
	var slice []int
	slice = append(slice, 1, 2, 3, 4)

	fmt.Println(reflect.TypeOf(slice))

	fmt.Println(reflect.TypeOf(arr))

	for _, val := range arr {
		fmt.Println(val)
	}

}
