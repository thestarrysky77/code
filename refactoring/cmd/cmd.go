package main

import (
	"fmt"

	"example.com/refactoring/pkg/expressions"
	"example.com/refactoring/pkg/function"
	"example.com/refactoring/pkg/object"
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

	fmt.Println("#####Extract Class#####")
	object.TestExtractClass()
	fmt.Println("#####Move Field#####")
	object.TestMoveField()
	fmt.Println("#####Move Method#####")
	object.TestMoveMethod()
	fmt.Println("#####Hide Delegate#####")
	object.TestHideDelegate()
	fmt.Println("#####Remove Middle Man#####")
	object.TestRemoveMiddleMan()
	fmt.Println("#####Local Extension#####")
	object.TestLocalExtension()

	fmt.Println("###Consolidate Conditional Expression###")
	expressions.TestConsolidateConditionalExpression()
	fmt.Println("###Consolidate Duplicate Conditional Fragments###")
	expressions.TestConsolidateDuplicateConditionalFragments()
	fmt.Println("###Introduce Null Object###")
	expressions.TestIntroduceNullObject()
	fmt.Println("###Remove Control Flag###")
	expressions.TestRemoveControlFlag()
	fmt.Println("###Replace Conditional With Polymorphism###")
	expressions.TestReplaceConditionalWithPolymorphism()
}
