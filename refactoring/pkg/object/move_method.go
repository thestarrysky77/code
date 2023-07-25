package object

import "fmt"

type AccountTypeOri1 struct {
}

func (accType AccountTypeOri1) isPremium() bool {
	return false
}

type AccountOri1 struct {
	AccountTypeOri1
	daysOverdrawn int
}

// --------------原始函数-----------------------
func (acc AccountOri1) bankCharge() float64 {
	result := 4.5
	if acc.daysOverdrawn > 0 {
		result += acc.overdraftCharge()
	}
	return result
}

func (acc AccountOri1) overdraftCharge() float64 {
	var result float64
	if acc.isPremium() {
		result = 10
		if acc.daysOverdrawn > 7 {
			result += (float64(acc.daysOverdrawn) - 7) * 0.05
		}
	} else {
		result = float64(acc.daysOverdrawn) * 1.75
	}
	return result
}

type AccountType1 struct {
}

func (accType AccountType1) isPremium() bool {
	return false
}

func (accType AccountType1) overdraftCharge(acc Account1) float64 {
	var result float64
	if accType.isPremium() {
		result = 10
		if acc.daysOverdrawn > 7 {
			result += (float64(acc.daysOverdrawn) - 7) * 0.05
		}
	} else {
		result = float64(acc.daysOverdrawn) * 1.75
	}
	return result
}

type Account1 struct {
	AccountType1
	daysOverdrawn int
}

func (acc Account1) bankCharge() float64 {
	result := 4.5
	if acc.daysOverdrawn > 0 {
		result += acc.overdraftCharge(acc)
	}
	return result
}

func TestMoveMethod() {
	accOri := AccountOri1{AccountTypeOri1{}, 15}
	fmt.Println("account charge: ", accOri.bankCharge())

	acc := Account1{AccountType1{}, 15}
	fmt.Println("account charge: ", acc.bankCharge())
}
