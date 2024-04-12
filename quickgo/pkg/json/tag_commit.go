package json

import (
	"encoding/json"
	"fmt"
)

// 测试两个struct序列化和反序列化

type From struct {
	Field1 int    `json:"field1,omitempty"`
	Field2 string `json:"field2,omitempty"`
	Field3 *int   `json:"field3,omitempty"`
}

type To struct {
	Field1 int    `json:"field1,omitempty"`
	Field2 string `json:"field2,omitempty"`
}

func testTagOverCommit() {
	i := 3
	f := From{1, "test", &i}
	fBytes, _ := json.Marshal(f)
	fmt.Println(string(fBytes))

	var t To
	json.Unmarshal(fBytes, &t)
	fmt.Println(t)
}

func testTagUnderCommit() {

	f := From{Field1: 1}
	fBytes, _ := json.Marshal(f)
	fmt.Println(string(fBytes))

	var t To
	json.Unmarshal(fBytes, &t)
	fmt.Println(t)
}
