package main

import (
	"fmt"

	"example.com/refactoring/src/pkg/function"
)

func main() {
	fmt.Println("#####Extract Method#####")
	function.TestExtractMethod()
	fmt.Println("#####Replace Temp With Query#####")
	function.TestReplaceTempWithQuery()
	fmt.Println("#####Introduce Explaining Variable#####")
	function.TestIntroduceExplainingVariable()
	fmt.Println("#####Split Temp Variable#####")
	function.TestSplitTempVariable()
	fmt.Println("#####Remove Assignments To Param#####")
	function.TestRemoveAssignmentsToParam()
	fmt.Println("#####Replace Method With Method Obj#####")
	function.TestReplaceMethodWithMethodObj()
}
