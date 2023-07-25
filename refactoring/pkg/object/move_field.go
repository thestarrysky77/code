package object

import "fmt"

type AccountOri struct {
	interestRate float64
}

// --------------原始函数-----------------------
func (acc AccountOri) interestForAmountDays(amount float64, days int) float64 {
	return acc.interestRate * amount * float64(days) / 365
}

type AccountType struct {
	interestRate float64
}

type Account struct {
	AccountType
}

func (acc Account) interestForAmountDays(amount float64, days int) float64 {
	return acc.interestRate * amount * float64(days) / 365
}

func TestMoveField() {
	accOri := AccountOri{0.04}
	fmt.Println("Interest: ", accOri.interestForAmountDays(50000, 180))

	acc := Account{AccountType{0.04}}
	fmt.Println("Interest: ", acc.interestForAmountDays(50000, 180))
}
