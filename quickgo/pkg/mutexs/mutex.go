package mutexs

import (
	"fmt"
	"sync"
)

var count int

func routineFunc(l *sync.RWMutex, c *sync.Cond, group *sync.WaitGroup) {
	defer group.Done()

	l.Lock()
	defer l.Unlock()

	c.Wait()
	fmt.Println("count: ", count)
}

func Produce(l *sync.RWMutex, c *sync.Cond, group *sync.WaitGroup) {
	defer group.Done()
	l.Lock()
	defer l.Unlock()

	// 测试signal和broadcast的区别。
	// signal 仅仅激活一个协程，broadcast激活所有协程
	// c.Signal()
	c.Broadcast()
	fmt.Println("signal")
}

func TestMutexCondition() {
	mutex := sync.RWMutex{}
	cond := sync.Cond{}
	cond.L = &mutex

	wg := sync.WaitGroup{}
	wg.Add(4)
	go routineFunc(&mutex, &cond, &wg)
	go routineFunc(&mutex, &cond, &wg)
	go routineFunc(&mutex, &cond, &wg)
	go Produce(&mutex, &cond, &wg)

	wg.Wait()
}
