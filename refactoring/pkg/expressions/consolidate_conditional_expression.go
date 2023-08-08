package expressions

import "fmt"

type Consolidate struct {
	seniority      int
	monthsDisabled int
	isPartTime     bool
}

// --- origin
func (c *Consolidate) disabilityAmountOri() int {
	if c.seniority < 2 {
		return 0
	}
	if c.monthsDisabled > 12 {
		return 0
	}
	if c.isPartTime {
		return 0
	}

	return 1
}

// --- modify
func (c *Consolidate) disabilityAmount() int {
	if c.isNotEligibleForDisability() {
		return 0
	}

	return 1
}

func (c *Consolidate) isNotEligibleForDisability() bool {
	return c.isPartTime || c.seniority < 2 || c.monthsDisabled > 12
}

// --- test
func TestConsolidateConditionalExpression() {
	c := new(Consolidate)
	c.isPartTime = true
	c.seniority = 1
	c.monthsDisabled = 11

	fmt.Println("disability amount ori: ", c.disabilityAmountOri())
	fmt.Println("disability amount: ", c.disabilityAmount())
}
