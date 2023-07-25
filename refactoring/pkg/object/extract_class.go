package object

import "fmt"

type PersonOri struct {
	name           string
	officeAreaCode string
	officeNumber   string
}

// --------------原始函数-----------------------
func (person PersonOri) getTelephoneNumber() string {
	return person.officeAreaCode + "-" + person.officeNumber
}

func (person PersonOri) print() {
	fmt.Println("person telephone: ", person.name, person.getTelephoneNumber())
}

type TelephoneNumber struct {
	officeAreaCode string
	officeNumber   string
}

type Person struct {
	TelephoneNumber
	name string
}

func (person Person) getTelephoneNumber() string {
	return person.officeAreaCode + "-" + person.officeNumber
}

func (person Person) print() {
	fmt.Println("person telephone: ", person.name, person.getTelephoneNumber())
}

func TestExtractClass() {
	personOri := PersonOri{"david", "0755", "86868686"}
	personOri.print()

	person := Person{TelephoneNumber{"0755", "5555555"}, "kitty"}
	person.print()
}
