package function

import "fmt"

type Order struct {
	amount float64
}

func (order Order) getAmount() float64 {
	return order.amount
}

type CustomOwes struct {
	name   string
	orders []Order
}

// --------------原始函数-----------------------
func (co CustomOwes) printOwingOri() {
	// print banner
	fmt.Println("**************************")
	fmt.Println("***** Customer Owes ******")
	fmt.Println("**************************")

	// calculate outstanding
	var outstanding float64
	for _, order := range co.orders {
		outstanding += order.getAmount()
	}

	//print details
	fmt.Println("name:", co.name)
	fmt.Println("amount", outstanding)
}

func (co *CustomOwes) AddOrder(order Order) {
	co.orders = append(co.orders, order)
}

func (co CustomOwes) printBanner() {
	fmt.Println("**************************")
	fmt.Println("***** Customer Owes ******")
	fmt.Println("**************************")
}

func (co CustomOwes) printDetails(outstanding float64) {
	fmt.Println("name:", co.name)
	fmt.Println("amount", outstanding)
}

func (co CustomOwes) calculateOutstanding() float64 {
	outstanding := 0.0
	for _, order := range co.orders {
		outstanding += order.getAmount()
	}

	return outstanding
}

func (co CustomOwes) printOwing() {
	co.printBanner()
	outstanding := co.calculateOutstanding()
	co.printDetails(outstanding)
}

func TestExtractMethod() {
	co := CustomOwes{name: "ayx"}
	co.AddOrder(Order{1})
	co.AddOrder(Order{2})
	co.AddOrder(Order{3})
	co.AddOrder(Order{4})

	co.printOwingOri()
	co.printOwing()
}
