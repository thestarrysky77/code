package channels

import (
	"fmt"
	"sync"
)

func routineFunc1(ch <-chan struct{}, group *sync.WaitGroup) {
	defer group.Done()

	_, ok := <-ch
	fmt.Println("Routine 1 state: ", ok)
}

func routineFunc2(ch <-chan struct{}, group *sync.WaitGroup) {
	defer group.Done()

	_, ok := <-ch
	fmt.Println("Routine 2 state: ", ok)
}

func TestChan() {
	stop := make(chan struct{}, 1)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go routineFunc1(stop, wg)
	go routineFunc2(stop, wg)

	close(stop)
	// stop <- struct{}{}

	wg.Wait()
}
