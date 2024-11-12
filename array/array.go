package main

import "fmt"

func main() {
	// var arr = [3] int {1,2,3}
	// fmt.Println(arr)

	// var nestedArr = [3][2] int {
	// 	{1,2},{3,4},{5,6},
	// }
	// fmt.Println(nestedArr)

	// var arr2 = [] int {4,5,6,7,8,9}
	// fmt.Println(arr2)
	// fmt.Println(len(arr2))
	// fmt.Println(cap(arr2))

	// arr3 := [] string {"hello","world","!!!"}
	// fmt.Println(arr3)

	// arr4 := [10] int {1,2,3,4,5,6,7,8,9,0}

	// // slice
	// slice := arr4[0:4]

	// fmt.Println(slice)

	var arr = [5]int{1, 2, 3}
	arr[4] = 4
	arr[3] = 5
	fmt.Println((arr))

	var dynArr = [...]int{1, 2, 3, 4, 5, 6}
	// dynArr[9] = 9
	fmt.Println(dynArr)

	var twoD [2][2]int

	twoD = [2][2]int{
		{1, 2}, {3, 4},
	}

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = twoD[i][j] + 5
		}
	}

	fmt.Println((twoD))
}
