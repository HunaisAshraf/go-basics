package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 1, 3, 4, 2, 5, 3, 9, 7}

	m := make(map[int]int)

	for _, val := range arr {

		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	// for _, v := range arr {
	// 	// fmt.Println(i, v)
	// 	m[v]++
	// }

	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)

	// m[1] = 145
	// m[2] = 2

	// fmt.Println((m))

	// value, exist := m[4]
	// fmt.Println(value, exist)

}
