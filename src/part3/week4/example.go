package main

import (
	"fmt"
	"sync"
)

func main() {

}

var on sync.Once

func setup() {
	fmt.Println("setup")
}

func dostuff() {
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}

var i int = 0
var mut sync.Mutex

func inc() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
}

/*func readChanel() {

	c := make(chan []int, 4)
	var a []int

	for {
		select {
		case a <- c:
			fmt.Println(a)

		case <-abot:
			return
		}
	}
}*/
