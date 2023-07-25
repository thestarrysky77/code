package object

import "fmt"

// --- origin
type PersonOri3 struct {
	name       string
	department DepartmentOri
}

type DepartmentOri struct {
	manager string
}

func (p *PersonOri3) GetManager() string {
	return p.department.GetManager()
}

func (p *PersonOri3) GetName() string {
	return p.name
}

func (d DepartmentOri) GetManager() string {
	return d.manager
}

// --- modify
type Person3 struct {
	name       string
	department DepartmentOri
}

func (p *Person3) GetDepartment() DepartmentOri {
	return p.department
}

func (p *Person3) GetName() string {
	return p.name
}

// --- test
func TestRemoveMiddleMan() {
	depart := DepartmentOri{"员工服务中心"}
	pOri3 := PersonOri3{name: "john", department: depart}
	fmt.Println("person manager: ", pOri3.GetName(), pOri3.GetManager())

	p3 := Person3{name: "john", department: depart}
	fmt.Println("person manager: ", p3.GetName(), p3.GetDepartment().GetManager())
}
