// 1. 从无缓存chan读：deadlock
// 2. 从有缓存有数据chan读：读成功
// 3. 从有缓存无数据chan读：deadlock
// 4. 从关闭chan读：false
// 5. 有发送者向接收chan发消息：读成功

// 总结：
// 1. 有数据、有发送者：读成功
// 2. 无数据：deadlock
// 3. 关闭：false

package channels

import (
	"fmt"
	"sync"
)

func readFromNocacheChan(ch <-chan struct{}, group *sync.WaitGroup) {
	defer group.Done()
	return

	value, ok := <-ch
	if ok {
		fmt.Println("value: ", value)
	}

	fmt.Println("read from no cache")
}

func readFromCacheChanWithData(ch <-chan int, group *sync.WaitGroup) {
	defer group.Done()

	value, ok := <-ch
	if ok {
		fmt.Println("value: ", value)
	}

	fmt.Println("read from cache chan with data")
}

func readFromCacheChanWithoutData(ch <-chan int, group *sync.WaitGroup) {
	defer group.Done()
	return

	value, ok := <-ch
	if ok {
		fmt.Println("value: ", value)
	}

	fmt.Println("read from cache chan without data")
}

func readFromClosedChan(ch <-chan int, group *sync.WaitGroup) {
	defer group.Done()

	value, ok := <-ch
	fmt.Println(value, ok)

	fmt.Println("read from closed chan")
}

func readFromChanWithSender(ch <-chan int, group *sync.WaitGroup) {
	defer group.Done()

	value, ok := <-ch
	if ok {
		fmt.Println("value: ", value)
	}

	fmt.Println("read from chan with sender")
}

func testReadChan() {
	noCacheChan := make(chan struct{})
	cacheChanWithData := make(chan int, 1)
	cacheChanWithData <- 1
	cacheChanWithoutData := make(chan int, 1)
	closedChan := make(chan int)
	close(closedChan)
	chanWithSender := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(5)
	go readFromNocacheChan(noCacheChan, wg)
	go readFromCacheChanWithData(cacheChanWithData, wg)
	go readFromCacheChanWithoutData(cacheChanWithoutData, wg)
	go readFromClosedChan(closedChan, wg)
	go readFromChanWithSender(chanWithSender, wg)
	chanWithSender <- 1

	wg.Wait()
}
