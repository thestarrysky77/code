package expressions

import "fmt"

func isSpecialDeal() bool {
	return true
}

func send() {
	fmt.Println("send a message")
}

// --- origin
func calPriceOri(price float64) float64 {
	var total float64

	if isSpecialDeal() {
		total = price * 0.95
		send()
	} else {
		total = price * 0.98
		send()
	}

	return total
}

// --- modify
func calPrice(price float64) float64 {
	var total float64

	if isSpecialDeal() {
		total = price * 0.95
	} else {
		total = price * 0.98
	}
	send()

	return total
}

func TestConsolidateDuplicateConditionalFragments() {
	price := 500
	fmt.Println("Price ori: ", calPriceOri(float64(price)))
	fmt.Println("Price: ", calPrice(float64(price)))
}
