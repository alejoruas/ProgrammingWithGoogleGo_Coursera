package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	id             int
	count          int
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

func (p Philosopher) eat(c chan *Philosopher, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p

		if p.count < 3 {
			p.leftChopStick.Lock()
			p.rightChopStick.Lock()
			fmt.Println("starting to eat ", p.id)
			p.count = p.count + 1
			time.Sleep(time.Second * 2)
			fmt.Println("finish eating ", p.id)
			p.rightChopStick.Unlock()
			p.leftChopStick.Unlock()
			wg.Done()
		}
	}
}

func host(c chan *Philosopher) {
	for {
		if len(c) == 2 {
			<-c
			<-c
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(15)

	c := make(chan *Philosopher, 2)

	ChopSticks := make([]*ChopStick, 5)

	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}

	Philosophers := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{
			i + 1,
			0,
			ChopSticks[i],
			ChopSticks[(i+1)%5],
		}
	}

	go host(c)

	for i := 0; i < 5; i++ {
		go Philosophers[i].eat(c, &wg)
	}

	wg.Wait()
}
