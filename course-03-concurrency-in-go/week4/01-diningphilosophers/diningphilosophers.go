package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	id              int
	mealsLeft       int
	name            string
}

var eatWgroup sync.WaitGroup
var currentlyEating int

var philo = [5]string{"Al-Kindi", "Al-Razi", "Al-Farabi", "Ibn Sina", "Ibn Rushd"}

func main() {

	//1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	cSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		//6. Each philosopher is numbered, 1 through 5.
		philos[i] = &Philo{cSticks[i], cSticks[(i+1)%5], i + 1, 3, philo[i]}
		eatWgroup.Add(1)
		go philos[i].Eat()
	}
	eatWgroup.Wait()

}

func (p Philo) Eat() {
	c := make(chan bool)

	//2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	for j := 0; j < 3; j++ {
		eatWgroup.Add(1)

		//4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		go getHostPermission(p, c)
		permitted := <-c

		if permitted {
			//3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
			p.leftCS.Lock()
			p.rightCS.Lock()

			//7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
			// fmt.Println(p.name, "starting to eat", p.id)
			fmt.Println("starting to eat", p.id)
			time.Sleep(time.Second)

			p.rightCS.Unlock()
			p.leftCS.Unlock()

			//8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
			p.mealsLeft--
			// fmt.Println(p.name, "finishing eating. He has", p.mealsLeft, " meals left")
			fmt.Println("finishing eating", p.id)
			currentlyEating--
			time.Sleep(time.Second)
			eatWgroup.Done()
		}
	}
	eatWgroup.Done()
}

func getHostPermission(p Philo, c chan bool) {

	for currentlyEating >= 2 {
		// fmt.Printf("Currently Eating: %d\n", currentlyEating)
		time.Sleep(time.Second)

	}

	if currentlyEating < 2 && p.mealsLeft > 0 {
		currentlyEating++
		//5. The host allows no more than 2 philosophers to eat concurrently.
		c <- true
	}

}
