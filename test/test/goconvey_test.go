package test

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return a / b, nil
}

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSub(t *testing.T) {
	Convey("将两数相减", t, func() {
		So(Sub(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相减乘", t, func() {
		So(Multiply(1, 2), ShouldEqual, 2)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {
		Convey("被除数为0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("被除数不为0", func() {
			num, err := Division(10, 2)
			So(num, ShouldEqual, 5)
			So(err, ShouldBeNil)
		})
	})
}

/**
每个单元测试的名称需要以 Test 开头，例如：TestAdd，并需要接受一个类型为 *testing.T 的参数。

使用 GoConvey 书写单元测试，每个测试用例需要使用 Convey 函数包裹起来。它接受的第一个参数为 string 类型的描述；第二个参数一般为 *testing.T，即本例中的变量 t；第三个参数为不接收任何参数也不返回任何值的函数（习惯以闭包的形式书写）。

Convey 语句同样可以无限嵌套，以体现各个测试用例之间的关系，例如 TestDivision 函数就采用了嵌套的方式体现它们之间的关系。需要注意的是，只有最外层的 Convey 需要传入变量 t，内层的嵌套均不需要传入。

作者：与蟒唯舞
链接：https://www.jianshu.com/p/1bd1ece2fa38
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
