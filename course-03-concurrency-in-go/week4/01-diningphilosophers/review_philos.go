package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i],
			CSticks[(i+1)%5], i}
	}
	// no more then two philos at the time
	h := &Host{seats: make(chan struct{}, 2)}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go h.serve(philos[i])
	}
	wg.Wait()
}

type ChopS struct{ sync.Mutex }
type Philo struct {
	leftCS, rightCS *ChopS
	idx             int
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("starting to eat ", p.idx)
		fmt.Println("finishing eating ", p.idx)
		p.rightCS.Unlock()
		p.leftCS.Unlock()

	}
	wg.Done()
}

type Host struct {
	seats chan struct{}
}

func (h *Host) serve(p *Philo) {
	// no more then two philos at the time
	h.seats <- struct{}{}
	p.eat()
	<-h.seats
}
