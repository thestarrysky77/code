package main

import "fmt"

type Abc struct {
	name string
}

type Test struct {
	name string
	abc  *[]Abc
}

func main() {
	test := Test{name: "test"}
	fmt.Println("len of abc:", len(*test.abc))
}
