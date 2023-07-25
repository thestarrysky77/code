package object

import (
	"fmt"
	"time"
)

// --- origin
type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) getDate() string {
	now := time.Now()
	d.year = now.Year()
	d.month = int(now.Month())
	d.day = now.Day()

	return fmt.Sprintf("%04d:%02d:%02d", d.year, d.month, d.day)
}

// --- modify
type ExDate struct {
	Date
}

func (d *ExDate) nextDay() string {
	now := time.Now()
	d.year = now.Year()
	d.month = int(now.Month())
	d.day = now.Day() + 1

	return fmt.Sprintf("%04d:%02d:%02d", d.year, d.month, d.day)
}

// --- test
func TestLocalExtension() {
	date := Date{}
	fmt.Println(date.getDate())

	exdate := ExDate{}
	fmt.Println(exdate.nextDay())
}
