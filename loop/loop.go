package main

import "fmt"

func main() {

	for i := range 4 {
		fmt.Println(i)
		if i == 2 {
			break
		}
	}

	n := 5

	if n%2 == 0 && n == 5 {
		fmt.Println("even")
	} else {
		fmt.Println(("odd"))
	}

}
