package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Id int `json:"id"`
}

func TestUnMarshal(t *testing.T) {

	body := `{"id":233322}`
	p := &Person{}
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("=====%v\n", p.Id)

}

/**
在Go语言中，json.Unmarshal函数用于将JSON数据解析为Go语言中的结构体或者接口类型。Go语言是一门静态类型语言，与Java不同，它强调强类型校验，因此在默认情况下，json.Unmarshal函数会进行强类型校验。

如果你希望在Go中像Java一样不进行强类型校验，可以使用json.RawMessage类型来延迟解析和处理JSON数据。json.RawMessage是一个字节切片类型，可以将任意的JSON数据存储在其中而不进行解析。
*/

type Person2 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Extra json.RawMessage
}

func TestUnMarshal2(t *testing.T) {

	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"extra": {
			"address": "123 Main St",
			"occupation": "Engineer"
		}
	}`

	var person Person2
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// 对Extra字段进行进一步解析
	var extraData map[string]interface{}
	err = json.Unmarshal(person.Extra, &extraData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Address:", extraData["address"])
	fmt.Println("Occupation:", extraData["occupation"])

}
