package main

import (
	"fmt"
	"sync"
)

type S struct {
	mut sync.Mutex
	m   map[string]int
}

func (s *S) add(i int) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.m["pos"] = i
}

// func something(id int, c <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for {
// 		val, ok := <-c
// 		fmt.Println(val, ok, id)
// 		if !ok {
// 			return
// 		}
// 	}
// }

func main() {

	var wg sync.WaitGroup

	// c := make(chan int)

	s := S{m: make(map[string]int)}

	// wg.Add(2)
	// go something(1, c, &wg)
	// go something(2, c, &wg)
	// go something(3, c, &wg)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.add(i)
		}(i)
	}

	// for i := 0; i < 30; i++ {
	// 	c <- i
	// }

	// close(c)

	wg.Wait()

	fmt.Println(s.m)

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(5 * time.Second)
	// 	c <- 5
	// 	fmt.Println("routine 1 done")
	// }()

	// output := <-c
	// fmt.Println(output)

}
