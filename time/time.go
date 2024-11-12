package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Tuesday)
	fmt.Println(time.Saturday)
	fmt.Println(time.Now().Hour())

	i := time.Now()
	switch {
	case i.Hour() < 10:
		fmt.Println("morning")
	case i.Hour() >= 10:
		fmt.Println(("office time"))
	}
}
