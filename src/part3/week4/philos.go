package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Each chopstick will act explicitly like a binary mutex flag, on to signify in use and off to signify available.

This will prevent multiple goroutines from using the same pair of chopsticks concurrently, because it will make a
goroutine block/wait if any of the chopsticks it attempts to use are already in use by another goroutine.

Since there are 5 chopsticks, this means that only 2 pairs (4 chopsticks) can be in use concurrently, the remaining 1
will have to block/wait for it's adjacent companion chopstick to free up and become available in order to use
it. So, only 2 philosophers at maximum can be eating concurrently, since each will need a pair of chopsticks
(2 chopsticks) and only 4 chopsticks can be in use concurrently.

BUT...

There is a possible interleaving in which all 5 philosophers pick up their left chopstick, or even their right
chopstick, all at the same time, rendering all of the chopsticks in use. This will create a deadlock because all 5
philosophers will be waiting to pick up their second chopstick forever.

One way of protecting against this is by allowing only 2 philosophers to pick up their pair of chopsticks at a time.
This is what the assignment suggests.
*/

/*
helpful articles:

http://adit.io/posts/2013-05-11-The-Dining-Philosophers-Problem-With-Ron-Swanson.html
https://en.wikipedia.org/wiki/Dining_philosophers_problem
*/

type Chop struct {
	sync.Mutex
}

type Philo struct {
	leftChop, rightChop *Chop
}

/*
helpful articles:

https://stackoverflow.com/questions/24868859/different-ways-to-pass-channels-as-arguments-in-function-in-go-golang
*/
func (p Philo) eat(pNum int, wg *sync.WaitGroup, host chan struct{}) {
	defer wg.Done()

	// eat only 3 times
	for i := 0; i < 3; i++ {
		// get permission from host to pick up your pair of chopsticks
		// send data to channel to signal request, and occupy 1 space in the buffer of 2
		/*
		helpful articles:

		https://medium.com/@l.peppoloni/how-to-improve-your-go-code-with-empty-structs-3bd0c66bc531
		https://mariadesouza.com/2019/01/12/empty-struct/
		https://programming.guide/go/broadcast-channel.html
		https://blog.golang.org/pipelines
		https://stackoverflow.com/questions/20793568/golang-anonymous-struct-and-empty-struct
		*/
		host <- struct{}{}

		// pick up your pair of chopsticks in any order, by generating a random number and testing its modulus for 1
		// (true) or 0 (false)
		/*
		helpful articles:

		https://stackoverflow.com/questions/49081879/selecting-a-random-case-in-switch-case
		https://stackoverflow.com/questions/44719156/how-can-i-let-a-function-randomly-return-either-a-true-or-a-false-in-go/44719269
		https://stackoverflow.com/questions/8393933/is-there-a-way-to-convert-integers-to-bools-in-go-or-vice-versa
		https://www.reddit.com/r/golang/comments/2vbtl4/int_as_bool/
		*/
		if rand.Int() % 2 == 1 {
			p.leftChop.Lock()
			p.rightChop.Lock()
		} else {
			p.rightChop.Lock()
			p.leftChop.Lock()
		}

		fmt.Println("Starting to eat", pNum + 1)

		fmt.Println("Finishing eating", pNum + 1)

		// put down your pair of chopsticks
		p.rightChop.Unlock()
		p.leftChop.Unlock()

		// tell host you are done with your pair of chopsticks
		// receive data from channel to signal request, and free up 1 space in the buffer of 2
		<- host
	}
}

func main() {
	// number of chopsticks and philosophers
	count := 5

	chops := make([]*Chop, count)

	// put chopsticks in slice of *Chop type
	for i := 0; i < count; i++ {
		chops[i] = new(Chop)
	}

	philos := make([]*Philo, count)

	// put philosophers in slice of *Philo type, and assign each the correct pair of chopsticks next to them
	for i := 0; i < count; i++ {
		philos[i] = &Philo{
			chops[i],
			chops[(i+1) % count],
		}
	}

	var wg sync.WaitGroup

	// use buffered channel to act as host, allowing only 2 philosophers to pick up their pair of chopsticks at a time
	host := make(chan struct{}, 2)

	for i := 0; i < count; i++  {
		wg.Add(1)

		// queue each philosopher to eat
		go philos[i].eat(i, &wg, host)
	}

	wg.Wait()
}
