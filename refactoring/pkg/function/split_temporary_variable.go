package function

import (
	"fmt"
	"math"
)

type Haggis struct {
	primaryForce   float64
	secondaryForce float64
	mass           float64
	delay          float64
}

// --------------原始函数-----------------------
func (ha Haggis) getDistanceTravelledOri(time int) float64 {
	var result float64

	acc := ha.primaryForce / ha.mass //第一次赋值处
	primaryTime := math.Min(float64(time), ha.delay)
	result = 0.5 * acc * primaryTime * primaryTime

	secondaryTime := float64(time) - ha.delay
	if secondaryTime > 0 {
		primaryVel := acc * ha.delay //以下是第二次赋值处
		acc = (ha.primaryForce + ha.secondaryForce) / ha.mass
		result += primaryVel*secondaryTime + 0.5*acc*secondaryTime*secondaryTime
	}

	fmt.Println("result: ", result)
	return result
}

// --------------原始函数-----------------------
func (ha Haggis) getDistanceTravelled(time int) float64 {
	var result float64

	primaryAcc := ha.primaryForce / ha.mass //第一次赋值处
	primaryTime := math.Min(float64(time), ha.delay)
	result = 0.5 * primaryAcc * primaryTime * primaryTime

	secondaryTime := float64(time) - ha.delay
	if secondaryTime > 0 {
		primaryVel := primaryAcc * ha.delay //以下是第二次赋值处
		secondaryAcc := (ha.primaryForce + ha.secondaryForce) / ha.mass
		result += primaryVel*secondaryTime + 0.5*secondaryAcc*secondaryTime*secondaryTime
	}

	fmt.Println("result: ", result)
	return result
}

func (ha Haggis) primaryAcc() float64 {
	return ha.primaryForce / ha.mass
}

func (ha Haggis) primaryTime(time float64) float64 {
	return math.Min(time, ha.delay)
}

func (ha Haggis) secondaryAcc() float64 {
	return (ha.primaryForce + ha.secondaryForce) / ha.mass
}

func (ha Haggis) secondaryTime(time float64) float64 {
	return time - ha.delay
}

func (ha Haggis) primaryDistance(time float64) float64 {
	return 0.5 * ha.primaryAcc() * ha.primaryTime(time) * ha.primaryTime(time)
}

func (ha Haggis) secondaryDistance(time float64) float64 {
	primaryVel := ha.primaryAcc() * ha.delay
	secondaryAcc := ha.secondaryAcc()
	return primaryVel*time + 0.5*secondaryAcc*time*time
}

func (ha Haggis) getDistanceTravelledFunc(time float64) float64 {
	var result float64

	result = ha.primaryDistance(time)

	if secondaryTime := ha.secondaryTime(time); secondaryTime > 0 {
		result += ha.secondaryDistance(secondaryTime)
	}

	fmt.Println("result: ", result)
	return result
}

func TestSplitTempVariable() {
	ha := Haggis{20, 30, 3, 6}
	ha.getDistanceTravelledOri(8)
	ha.getDistanceTravelled(8)
	ha.getDistanceTravelledFunc(8)
}
