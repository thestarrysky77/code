package function

import "fmt"

type Counter struct {
	quantity  int64
	itemPrice float64

	discountFactor float64
}

// --------------原始函数-----------------------
func (c Counter) getPriceOri() float64 {
	var basePrice float64
	basePrice = float64(c.quantity) * c.itemPrice

	var discountFactor float64
	if basePrice > 1000 {
		discountFactor = 0.95
	} else {
		discountFactor = 0.98
	}

	return basePrice * discountFactor
}

func (c Counter) getBasePrice() float64 {
	return float64(c.quantity) * c.itemPrice
}

func (c Counter) getDiscountFactor() float64 {
	var discountFactor float64
	if c.getBasePrice() > 1000 {
		discountFactor = 0.95
	} else {
		discountFactor = 0.98
	}

	return discountFactor
}

func (c Counter) getPrice() float64 {
	return c.getBasePrice() * c.getDiscountFactor()
}

func TestReplaceTempWithQuery() {
	c := Counter{quantity: 200, itemPrice: 6.5}
	fmt.Println("origin: ", c.getPriceOri())
	fmt.Println("replace temp with query: ", c.getPrice())
}
