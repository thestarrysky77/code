package expressions

import "fmt"

type Object struct {
	name string
}

func (obj *Object) getName() string {
	return obj.name
}

// --- origin
func getObjectOri(name string) *Object {
	if name == "" {
		return nil
	}

	obj := Object{name: name}

	return &obj
}

func getNameOri() string {
	obj := getObjectOri("")
	if obj == nil {
		return "nil"
	}
	return obj.getName()
}

// --- modify
func getObject(name string) *Object {
	obj := Object{}
	if name == "" {
		obj.name = "nil"
	} else {
		obj.name = name
	}
	return &obj
}

func getName() string {
	obj := getObject("")
	return obj.getName()
}

func TestIntroduceNullObject() {
	fmt.Println("Get Name Ori: ", getNameOri())
	fmt.Println("Get Name: ", getName())
}
