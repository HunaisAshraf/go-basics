package main

import "fmt"

var something string = "something";


func main(){
	a := 123
	something = "changed"
	fmt.Println(something)
	fmt.Println(a)

	var num1 = 10;
	var num2 int = 20;
	var num3 uint = 30;

	num4 := 60

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	var name string = "hello"
	name2 := "world"

	fmt.Println(name)
	fmt.Println(name2)

	var name3 string
	name3 = "!!!!!!!"

	fmt.Println(name3)

	// grouping

	var one, two, three = 1,2,3
	fmt.Println(one,two,three)
	four,five,six := 4,5,6
	fmt.Println(four, five, six)

	var (
		eight=8
		nine=9
		ten=10
	)

	fmt.Println(eight,nine,ten)

	const NAME string = "hunais"
	fmt.Println(NAME)

	fmt.Print(four,five,six)
	fmt.Print(name," ",NAME, "\n")

	fmt.Printf("value = %#v type = %T", NAME,NAME)

}