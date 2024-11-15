package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// go runLoop("hello")
	// runLoop("world")

	// sites := []string{
	// 	"http://www.google.com",
	// 	"http://www.aech.online",
	// 	"http://www.google.com",
	// 	"http://www.aech.online",
	// 	"http://www.google.com",
	// 	"http://www.aech.online",
	// 	"http://www.google.com",
	// 	"http://www.aech.online",
	// 	"http://www.google.com",
	// 	"http://www.aech.online",
	// 	"http://www.google.com",
	// 	"http://www.aech.online",
	// 	"http://www.google.com",
	// 	"http://www.aech.online",
	// }

	// func something(){
	//
	// }
	//

	// for _, site := range sites {
	// 	// fmt.Println(site)
	// 	go fetchData(site)
	// 	wg.Add(1)
	// }

	// wg.Wait()

}

func fetchData(site string) {
	defer wg.Done()
	res, err := http.Get(site)

	if err != nil {
		panic(err)

		// fmt.Println(err)
	} else {
		fmt.Printf("%v status is %v\n", site, res.StatusCode)
	}

	// jsonData := string(data)
	// fmt.Println(jsonData)

	defer res.Body.Close()
}

func runLoop(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(str)
	}
}
