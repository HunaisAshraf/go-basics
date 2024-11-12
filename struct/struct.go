package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
	phone int
}

func main() {
	var p1 Person

	p1.name = "aaa"
	p1.age = 44
	p1.email = "something@gmail.com"
	p1.phone = 123456789

	fmt.Println(p1)

	p2 := Person{name: "nnnn", age: 55, email: "alksjdf@gmail.com", phone: 98734298493}
	fmt.Println(p2)

	p3 := p2
	p3.name = "zzzzz"
	fmt.Println(p2)
	fmt.Println(p3)

	p4 := createPerson("aklsdjf", 23, "kjsdfasfdsadf", 53465465)
	fmt.Println(p4)

}

func createPerson(name string, age int, email string, phone int) *Person {
	p := Person{name, age, email, phone}
	return &p
}
