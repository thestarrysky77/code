package main

import (
	"example.com/quickgo/pkg/channels"
	"example.com/quickgo/pkg/defaults"
	"example.com/quickgo/pkg/json"
	"example.com/quickgo/pkg/mutexs"
)

func main() {
	// defaults
	defaults.TestAll()

	// json
	json.TestAll()

	// channels
	channels.TestChan()

	// mutexs
	mutexs.TestMutexCondition()
}
