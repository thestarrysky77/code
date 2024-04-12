package defaults

import "fmt"

func testInteger() {
	var i int
	fmt.Println(i)
}

func testFloat() {
	var f float64
	fmt.Println(f)
}

func testString() {
	var s string
	fmt.Println(s)
}

func testBool() {
	var b bool
	fmt.Println(b)
}

func testArray() {
	var ai [3]int
	fmt.Println(ai)
}

func testStruct() {
	type S struct {
		i  int
		pi *int
		s  []string
	}
	var s S
	fmt.Println(s)
}
