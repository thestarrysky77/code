package json

import (
	"encoding/json"
	"fmt"
)

type OmitStuct struct {
	Field1 string   `json:"field1,omitempty"`
	Field2 *string  `json:"field2,omitempty"`
	Field3 []string `json:"field3,omitempty"`
	Field4 int      `json:"field4"`
	Field5 *int     `json:"field5"`
}

// testOmit omitempty字段转为json串时0值将被忽略
func testOmit() {

	// JSON编码（序列化）时：
	// omitempty字段0值将被忽略
	// 非omitempty字段0值将被输出
	o1 := OmitStuct{Field1: "test"}
	o1Bytes, _ := json.Marshal(o1)
	fmt.Println(string(o1Bytes))
}
