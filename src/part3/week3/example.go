package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	fmt.Printf("New routine")
	wg.Done()
}

/*func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Printf("Main routine")
}*/

func prod(v1 int, v2 int, c chan int) {
	c <- v1 * v2
}

func main() {
	c := make(chan int)
	go prod(1, 2, c)
	go prod(3, 4, c)
	a := <-c
	b := <-c
	fmt.Println(a * b)
}
