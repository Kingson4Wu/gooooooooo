package test

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}

/**
这个测试用例有2个目的，一是使用 ctrl.Finish() 断言 DB.Get() 被是否被调用，如果没有被调用，后续的 mock 就失去了意义；
二是测试方法 GetFromDB() 的逻辑是否正确(如果 DB.Get() 返回 error，那么 GetFromDB() 返回 -1)。
NewMockDB() 的定义在 db_mock.go 中，由 mockgen 自动生成。

*/

/**
在上面的例子中，当 Get() 的参数为 Tom，则返回 error，这称之为打桩(stub)，有明确的参数和返回值是最简单打桩方式。除此之外，检测调用次数、调用顺序，动态设置返回值等方式也经常使用。
*/

/**
写可测试的代码与写好测试用例是同等重要的，如何写可 mock 的代码呢？

mock 作用的是接口，因此将依赖抽象为接口，而不是直接依赖具体的类。
不直接依赖的实例，而是使用依赖注入降低耦合性。
*/

/**
如果 GetFromDB() 方法长这个样子

func GetFromDB(key string) int {
	db := NewDB()
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
对 DB 接口的 mock 并不能作用于 GetFromDB() 内部，这样写是没办法进行测试的。那如果将接口 db DB 通过参数传递到 GetFromDB()，那么就可以轻而易举地传入 Mock 对象了。

*/
