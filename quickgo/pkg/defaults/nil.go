package defaults

import "fmt"

func testPointer() {
	var pi *int
	fmt.Println(pi)
}

func testSlice() {
	var is []int
	fmt.Println(is)
}

func testMap() {
	var m map[int]int
	fmt.Println(m)
}

func testChan() {
	var c chan int
	fmt.Println(c)
}

func testFunc() {
	type Func func()
	var fun Func
	fmt.Println(fun)
}

func testInterface() {
	var i interface{}
	fmt.Println(i)
}
