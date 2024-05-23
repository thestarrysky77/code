package main

import (
	"example.com/quickgo/pkg/channels"
	"example.com/quickgo/pkg/defaults"
	"example.com/quickgo/pkg/json"
)

func main() {
	// defaults
	defaults.TestAll()

	// json
	json.TestAll()

	// channels
	channels.TestAll()

	// mutexs
	//mutexs.TestMutexCondition()
}
