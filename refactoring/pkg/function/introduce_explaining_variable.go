package function

import (
	"fmt"
	"math"
)

type Price struct {
	quantity  int64
	itemPrice float64
}

// --------------原始函数-----------------------
func (p Price) priceOri() float64 {
	// price is base price - quantity discount + shipping
	basePrice := float64(p.quantity) * p.itemPrice
	price := basePrice -
		math.Max(0, float64(p.quantity)-500)*p.itemPrice*0.05 +
		math.Min(float64(p.quantity)*p.itemPrice*0.1, 100)
	fmt.Println("price: ", price)
	return price
}

func (p Price) price() float64 {
	basePrice := float64(p.quantity) * p.itemPrice
	discount := math.Max(0, float64(p.quantity)-500) * p.itemPrice * 0.05
	shipping := math.Min(float64(p.quantity)*p.itemPrice*0.1, 100)
	price := basePrice - discount + shipping
	fmt.Println("price: ", price)
	return price
}

func (p Price) basePrice() float64 {
	return float64(p.quantity) * p.itemPrice
}

func (p Price) discount() float64 {
	return math.Max(0, float64(p.quantity)-500) * p.itemPrice * 0.05
}

func (p Price) shipping() float64 {
	return math.Min(float64(p.quantity)*p.itemPrice*0.1, 100)
}

func (p Price) priceFunc() float64 {
	price := p.basePrice() - p.discount() + p.shipping()
	fmt.Println("price: ", price)
	return price
}

func TestIntroduceExplainingVariable() {
	p := Price{400, 5.6}
	p.priceOri()
	p.price()
	p.priceFunc()
}
