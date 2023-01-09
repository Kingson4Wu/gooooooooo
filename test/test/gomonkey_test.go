package test

import (
	"encoding/json"
	reflect "reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
)

// mock一个函数 -----

func networkCompute(a, b int) (int, error) {
	// do something in remote computer
	c := a + b

	return c, nil
}

func Compute(a, b int) (int, error) {
	sum, err := networkCompute(a, b)
	return sum, err
}

func TestCompute(t *testing.T) {

	patches := gomonkey.ApplyFunc(networkCompute, func(a, b int) (int, error) {
		return 2, nil
	})

	defer patches.Reset()

	sum, err := Compute(1, 1)
	if sum != 2 || err != nil {
		t.Errorf("expected %v, got %v", 2, sum)
	}

}

//----

type Host struct {
	IP   string
	Name string
}

func Convert2Json(h *Host) (string, error) {
	b, err := json.Marshal(h)
	return string(b), err
}

func TestConvert2Json(t *testing.T) {
	patches := gomonkey.ApplyFunc(json.Marshal, func(v interface{}) ([]byte, error) {
		return []byte(`{"IP":"192.168.23.92","Name":"Sky"}`), nil
	})

	defer patches.Reset()

	h := Host{Name: "Sky", IP: "192.168.23.91"}
	s, err := Convert2Json(&h)

	expectedString := `{"IP":"192.168.23.92","Name":"Sky"}`

	if s != expectedString || err != nil {
		t.Errorf("expected %v, got %v", expectedString, s)
	}

}

//mock 一个方法 ------

type Computer struct {
}

func (t *Computer) NetworkCompute(a, b int) (int, error) {
	// do something in remote computer
	c := a + b

	return c, nil
}

func (t *Computer) Compute(a, b int) (int, error) {
	sum, err := t.NetworkCompute(a, b)
	return sum, err
}

func TestCompute2(t *testing.T) {
	var c *Computer
	patches := gomonkey.ApplyMethod(reflect.TypeOf(c), "NetworkCompute", func(_ *Computer, a, b int) (int, error) {
		return 2, nil
	})

	defer patches.Reset()

	cp := &Computer{}
	sum, err := cp.Compute(1, 1)
	if sum != 2 || err != nil {
		t.Errorf("expected %v, got %v", 2, sum)
	}

}

//mock 一个全局变量 -----

var num = 10

func TestGlobalVar(t *testing.T) {
	patches := gomonkey.ApplyGlobalVar(&num, 12)
	defer patches.Reset()

	if num != 12 {
		t.Errorf("expected %v, got %v", 12, num)
	}
}

//mock 一个函数序列 -----
//函数序列主要用在，一个函数被多次调用，每次调用返回不同值。

func compute(a, b int) (int, error) {
	return a + b, nil
}

func TestFunc(t *testing.T) {
	info1 := "2"
	info2 := "3"
	info3 := "4"
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{info1, nil}}, // 模拟函数的第1次输出
		{Values: gomonkey.Params{info2, nil}}, // 模拟函数的第2次输出
		{Values: gomonkey.Params{info3, nil}}, // 模拟函数的第3次输出
	}
	patches := gomonkey.ApplyFuncSeq(compute, outputs)
	defer patches.Reset()

	output, err := compute(1, 1)
	if output != 2 || err != nil {
		t.Errorf("expected %v, got %v", 2, output)
	}

	output, err = compute(1, 2)
	if output != 3 || err != nil {
		t.Errorf("expected %v, got %v", 2, output)
	}

	output, err = compute(1, 3)
	if output != 4 || err != nil {
		t.Errorf("expected %v, got %v", 2, output)
	}

}

/**
有时会遇到mock失效的情况，这个问题一般是内联导致的。

什么是内联？

为了减少函数调用时的堆栈等开销，对于简短的函数，会在编译时，直接内嵌调用的代码。
*/

var flag bool

func IsEnabled() bool {
	return flag
}

func Compute3(a, b int) int {
	if IsEnabled() {
		return a + b
	}

	return a - b
}

func TestCompute3(t *testing.T) {
	patches := gomonkey.ApplyFunc(IsEnabled, func() bool {
		return true
	})

	defer patches.Reset()

	sum := Compute3(1, 1)
	if sum != 2 {
		t.Errorf("expected %v, got %v", 2, sum)
	}

}

/**
关闭内联的，再次尝试：

go test -v -gcflags=-l gomonkey_test.go
*/
