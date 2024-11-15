package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name  string `json:"userName"`
	Email string `json:"email"`
	Age   int    `json:"age,omitempty"`
}

func main() {
	encodeJson()
}

func encodeJson() {
	jsonData := []Data{
		{Name: "someName", Email: "someEmail", Age: 44},
		{Name: "name2", Email: "email2", Age: 33},
		{Name: "name3", Email: "email3", Age: 22},
		{Name: "name3", Email: "email3"},
	}

	fmt.Println(jsonData)

	finalData, err := json.MarshalIndent(jsonData, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalData)
}
