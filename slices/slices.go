package main

import "fmt"

func main() {
	// var newSlice = make([]int, 5,10)
	// fmt.Println(len(newSlice),cap(newSlice),newSlice)
	// newSlice[3] = 8

	// var slice []int

	slice := make([]int, 5, 10)
	slice[4] = 8
	slice = append(slice, 10, 40)
	// slice = append(slice, 40)
	fmt.Println(len(slice), cap(slice), slice)

	newSlice := make([]int, 10, 10)
	copy(newSlice, slice)
	fmt.Println(newSlice)

	anotherSLice := make([]int, 10, 10)
	copy(anotherSLice, newSlice)
	newSlice[4] = 1000
	fmt.Println(newSlice)
	fmt.Println(anotherSLice)

	slicing := anotherSLice[4:7]

	fmt.Println(slicing)

	s1 := [][]int{{1, 2, 3}, {1, 2, 3}}

	fmt.Println(s1)

}
