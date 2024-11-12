package main

import (
	"fmt"
	"strings"
)

func main() {

	const str = "hello world"

	fmt.Println(len(str))

	for i, val := range str {
		fmt.Println(i, val)
	}

	arr := strings.Split(str, " ")
	fmt.Println(arr)

	for _, val := range arr {
		fmt.Printf("value = %v, type = %T \n", val, val)
	}
}
