package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(div(1, 2))

	_, b := multRetrun()

	fmt.Println(b)
	fmt.Println(returnArr())

	fmt.Println(variadicFunction(1, 2, 3, 4, 5, 6))

	cl := closure()

	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(fun())

	fmt.Println(fact(5))
}

func add(a int, b int) int {
	return a + b
}

func div(a, b int) int {
	return a / b
}

func multRetrun() (int, string) {
	return 1, "2"
}

func returnArr() [5]int {
	arr := [5]int{1, 2, 3, 4, 5}

	return arr
}

func variadicFunction(a, b int, nums ...int) int {
	fmt.Println(reflect.TypeOf(nums))
	fmt.Println(a, b, nums)

	total := 0

	for _, val := range nums {
		total += val
	}

	fmt.Println(total)

	return a + b + total

}

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var fun = func() string {
	return "hello"
}

func fact(n int) int {
	if n == 1 {
		return n
	}

	return fact(n-1) * n
}
