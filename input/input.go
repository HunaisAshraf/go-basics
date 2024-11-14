package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter something")
	input, _ := reader.ReadString('\n')
	fmt.Println("entered input is: ", input)
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num + 100)
}
