package main

import "fmt"

func main() {
	i := 1
	value(i)
	fmt.Println(i)
	fmt.Println(&i)
	pointer(&i)
	fmt.Println(i)
	fmt.Println(&i)
	value(i)
	fmt.Println(i)
}

func pointer(val *int) {
	*val = 0
}

func value(val int) {
	val = 0
}
