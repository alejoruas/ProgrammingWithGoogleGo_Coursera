package main

import (
	"time"
)

var num int

func count() {
	num++
}

func main() {
	go count()
	go count()
	time.Sleep(3 * time.Second)
}
