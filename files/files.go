package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "new file is created and added some value to it"

	file, err := os.Create("./newFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length of the file is", length)
	defer file.Close()
	reader("./newFile.txt")
}

func reader(fileName string) string {
	data, err := os.ReadFile(fileName)
	checkNilError(err)
	return string(data)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
