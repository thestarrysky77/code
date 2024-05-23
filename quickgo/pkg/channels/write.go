// 1. 向无缓存chan写数据：deadlock
// 2. 向有缓存，有空间chan写数据：成功
// 3. 向有缓存，无空间chan写数据：deadlock
// 4. 向关闭的chan写数据：panic
// 5. 向有接收者的chan写数据：成功，接收者被激活

// 总结：
// 1. chan有空间、有接受者写成功
// 2. chan无空间，写deadlock
// 3. chan关闭，写panic

package channels

import (
	"fmt"
	"sync"
)

func writeToNoCacheChan(group *sync.WaitGroup) {
	defer group.Done()
	return

	noCacheChan := make(chan int, 0)
	noCacheChan <- 1
	fmt.Println("write to no cache chan")
}

func writeToCacheChanWithSpace(group *sync.WaitGroup) {
	defer group.Done()
	return

	cacheChan := make(chan int, 1)
	cacheChan <- 1
	fmt.Println("write to cache chan with space")
}

func writeToCacheChanWithoutSpace(group *sync.WaitGroup) {
	defer group.Done()
	return

	cacheChan := make(chan int, 1)
	cacheChan <- 1
	cacheChan <- 2
	fmt.Println("write to cache chan without space")
}

func writeToClosedChan(group *sync.WaitGroup) {
	defer group.Done()
	return

	cacheChan := make(chan int, 1)
	close(cacheChan)
	cacheChan <- 1
	fmt.Println("write to closed chan")
}

func writeToChanWithReceiver(group *sync.WaitGroup) {
	defer group.Done()

	wg := &sync.WaitGroup{}
	ch := make(chan int, 0)
	wg.Add(1)
	go receiver(ch, wg)
	ch <- 1
	fmt.Println("write to chan with receiver")
	wg.Wait()
}

func receiver(ch <-chan int, group *sync.WaitGroup) {
	defer group.Done()
	value, ok := <-ch
	if ok {
		fmt.Println("value: ", value)
	}
	fmt.Println("receiver")

}

func testWriteChan() {
	g := &sync.WaitGroup{}
	g.Add(5)
	go writeToNoCacheChan(g)
	go writeToCacheChanWithSpace(g)
	go writeToCacheChanWithoutSpace(g)
	go writeToClosedChan(g)
	go writeToChanWithReceiver(g)

	g.Wait()
}
