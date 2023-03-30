package function

import "fmt"

// --------------原始函数-----------------------
func discountOri(inputVal int, quantity int, yearToDate int) int {
	if inputVal > 50 {
		inputVal -= 2
	}
	if quantity > 100 {
		inputVal -= 1
	}
	if yearToDate > 10000 {
		inputVal -= 4
	}
	fmt.Println("input value: ", inputVal)
	return inputVal
}

func discount(inputVal int, quantity int, yearToDate int) int {
	result := inputVal
	if inputVal > 50 {
		result -= 2
	}
	if quantity > 100 {
		inputVal -= 1
	}
	if yearToDate > 10000 {
		inputVal -= 4
	}
	fmt.Println("input value: ", result)
	return result
}

func TestRemoveAssignmentsToParam() {
	discountOri(100, 3, 300)
	discount(100, 3, 300)
}
