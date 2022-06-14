package main

import (
	"fmt"
	"runtime"
	"sync"
)

var x = 0

/*
There is a race condiktion abput de the increment of X because of concurrent access of the go-rutines over x.
This occur due to access to X at the same time before de increment of every go-rutine.
This was tested with 100 go-rutines trying to increment x*/

func increase(wg *sync.WaitGroup) {
	temp := x
	runtime.Gosched()
	temp++
	x = temp
	wg.Done()
}

func main() {

	const num = 100

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go increase(&wg)
	}
	wg.Wait()
	fmt.Println("x: ", x)
}
