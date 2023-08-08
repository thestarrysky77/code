package expressions

import "fmt"

type Pay interface {
	PayAmount() float64
}

type EmployeeType struct {
	name       string
	baseSalary float64
}

func (e *EmployeeType) getBaseSalary() float64 {
	return e.baseSalary
}

func (e *EmployeeType) PayAmount() float64 {
	return e.getBaseSalary()
}

type Engineer struct {
	EmployeeType
}

type Manager struct {
	EmployeeType
	bonus float64
}

func (e *Manager) PayAmount() float64 {
	return e.getBaseSalary() + e.bonus
}

type SalesMan struct {
	EmployeeType
	commission float64
}

func (e *SalesMan) PayAmount() float64 {
	return e.getBaseSalary() + e.commission
}

func TestReplaceConditionalWithPolymorphism() {

	Employees := []Pay{
		&Engineer{
			EmployeeType{
				name:       "e",
				baseSalary: 6000.50,
			},
		},
		&Manager{
			EmployeeType: EmployeeType{
				name:       "e",
				baseSalary: 6000.50,
			},
			bonus: 7000,
		},
		&SalesMan{
			EmployeeType: EmployeeType{
				name:       "e",
				baseSalary: 6000.50,
			},
			commission: 3000,
		},
	}

	for _, e := range Employees {
		fmt.Println("PayAmount: ", e.PayAmount())
	}
}
