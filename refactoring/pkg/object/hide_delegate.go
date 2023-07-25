package object

import "fmt"

// --- origin
type PersonOri2 struct {
	name       string
	department Department
}

type Department struct {
	manager string
}

func (p *PersonOri2) GetDepartment() Department {
	return p.department
}

func (p *PersonOri2) GetName() string {
	return p.name
}

func (d Department) GetManager() string {
	return d.manager
}

// --- modify
// --- origin
type Person2 struct {
	name       string
	department Department
}

func (p *Person2) GetName() string {
	return p.name
}

func (p *Person2) GetManager() string {
	return p.department.GetManager()
}

// --- test
func TestHideDelegate() {
	depart := Department{"员工服务中心"}
	pOri2 := PersonOri2{name: "john", department: depart}
	fmt.Println("person manager: ", pOri2.GetName(), pOri2.GetDepartment().GetManager())

	p2 := Person2{name: "john", department: depart}
	fmt.Println("person manager: ", p2.GetName(), p2.GetManager())
}
