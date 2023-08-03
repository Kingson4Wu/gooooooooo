package errrr

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type Err struct {
	err string
}

func (e *Err) Error() string {
	return e.err
}

func returnErr() *Err {
	return nil
}

func TestError(t *testing.T) {

	var err error
	err = returnErr()
	fmt.Printf("err: %+v, compare: %+v\n", err, err == nil)
}

func TestError2(t *testing.T) {

	err := returnErr()
	fmt.Printf("err: %+v, compare: %+v\n", err, err == nil)
}

var ErrNil *Err

func TestError3(t *testing.T) {

	var err error
	err = returnErr()
	fmt.Printf("err: %+v, compare: %+v\n", err, err == ErrNil)
}

func TestErrorFormat(t *testing.T) {
	err := errors.New("panic")
	//err = nil
	fmt.Printf("===: %s\n", err)
	fmt.Printf("===: %v\n", err)
	fmt.Printf("===: %+v\n", err)
	fmt.Printf("===: %s\n", err)
	fmt.Printf("===: %s\n", err)
}

func Test1(t *testing.T) {
	var nn interface{}
	dd, _ := nn.(string)
	fmt.Println(dd)

	contentType, _ := nn.(map[string]interface{})["contenttype"].(int)
	fmt.Println(contentType)
}

type Parser struct {
	name  string
	parse func(string)
}

var parserList []*Parser

func AddParser(name string, parse func(string)) {
	p := &Parser{name: name, parse: parse}
	//p.parse(appConfig.GetString(name, ""))
	parserList = append(parserList, p)
}

type person struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func TestChangeListen(t *testing.T) {

	var m map[string]string
	var p *person

	AddParser("layu", func(s string) {
		var mm map[string]string
		json.Unmarshal([]byte(s), &mm)
		m = mm
	})

	AddParser("layu", func(s string) {
		var pp person
		json.Unmarshal([]byte(s), &pp)
		p = &pp
	})

	fmt.Println(m)
	fmt.Println(p)

	for _, p := range parserList {
		newValue := `{"name":"torres", "desc":"goooo"}`
		p.parse(newValue)
	}

	fmt.Println(m)
	fmt.Println(p)

}
