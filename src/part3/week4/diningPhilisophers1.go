package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
- Each philosopher should eat only 3 times
- The philosophers pick up the chopsticks in any order, not lowest-numbered first.
- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
- The host allows no more than 2 philosophers to eat concurrently.
- Each philosopher is numbered, 1 through 5.
- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

var wgHappyPhilosophers = sync.WaitGroup{}

type ChopS struct {
	sync.Mutex
}

type Philosopher struct {
	number          int
	leftCS, rightCS *ChopS
	eatTimes        int
}

type Host struct {
	eatAtOneTime int
}

func (h *Host) rule(requestToEat chan int, permissionToEat chan<- int, doneEating <-chan int, abort <-chan interface{}) {
	currentEating := 0
	for{
		select{
		case p := <- requestToEat:
			fmt.Printf("\nHost <- request from %v", p)
			if currentEating < h.eatAtOneTime{
				currentEating++
				permissionToEat <- 0
				fmt.Printf("\nHost permission given <-, curently %v philosopher eating", currentEating)
			} else {
				requestToEat <- p
			}
		case p := <- doneEating:{
			fmt.Printf("\nHost <- done eating %v", p)
			currentEating--
		}
		case <- abort:
			return
		}
	}
}

func (p *Philosopher) eat(requestToEat chan<- int, permissionToEat <-chan int, doneEating chan<- int) {
	for {
		if p.eatTimes < 3 {

			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("\nPhilosopher %v obtained sticks", p.number)

			requestToEat <- p.number

			<- permissionToEat

			p.eatTimes++
			fmt.Printf("\nPhilosopher %v starting to eat, %v time", p.number, p.eatTimes)
			time.Sleep(10000)

			doneEating <- p.number
			fmt.Printf("\nPhilosopher %v finishing eating, %v time", p.number, p.eatTimes)

			p.rightCS.Unlock()
			p.leftCS.Unlock()

			// Some rest time :)
			time.Sleep(1000)

		} else {
			fmt.Printf("\nPhilisopher %v Done!", p.number)
			wgHappyPhilosophers.Done()
			break
		}
	}
}

func main() {
	fmt.Printf("\n######## Concurrency in Go. Week 4. Dining Philisophers. #########\n")

	wgHappyPhilosophers.Add(5)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5], 0}
	}

	requestToEat := make(chan int)
	permissionToEat := make(chan int)
	doneEating := make(chan int)
	abort := make(chan interface{})

	host := Host{2}

	go host.rule(requestToEat, permissionToEat, doneEating, abort)

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(requestToEat, permissionToEat, doneEating)
	}

	wgHappyPhilosophers.Wait()
	abort <- true
}
