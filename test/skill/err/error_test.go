package errrr

import (
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
	fmt.Printf("err: %+v, compare: %+v\n", err, err == nil) //false
}

func TestError2(t *testing.T) {

	err := returnErr()
	fmt.Printf("err: %+v, compare: %+v\n", err, err == nil) // true
}

/**
在第一个测试用例 TestError 中，err 是一个 error 类型的变量，它是一个接口类型。
当你将 returnErr() 的返回值赋给 err 时，实际上是将一个 nil 指针赋给了 err 变量。
然而，在比较 err 是否为 nil 时，你比较的是 err 接口类型变量本身是否为 nil，而不是它包含的错误值。
由于 err 是一个接口类型，它并不会直接比较其包含的错误值是否为 nil，而是比较接口本身是否为 nil。
所以在这个情况下，err == nil 返回 false。

在第二个测试用例 TestError2 中，err 是一个具体的错误类型 *Err 的变量，而不是接口类型。
当你将 returnErr() 的返回值赋给 err 时，你实际上将一个 nil 指针赋给了 *Err 类型的变量，
因此 err 本身就是 nil。因此，当你比较 err 是否为 nil 时，它会返回 true。
*/

var ErrNil *Err

func TestError3(t *testing.T) {

	var err error
	err = returnErr()
	fmt.Printf("err: %+v, compare: %+v\n", err, err == ErrNil) //true
}

/**
在这个测试用例中，你声明了一个全局变量 ErrNil，并将其初始化为 *Err 类型的 nil 指针。
然后，在 TestError3 中，你将 returnErr() 的返回值赋给了 err 变量，
并且将 err 与 ErrNil 进行了比较。

因为 err 是一个 error 类型的变量，而 ErrNil 是一个 *Err 类型的指针，它们的类型不匹配，
所以直接使用 err == ErrNil 进行比较会导致编译错误，因为 Go 不允许不同类型的变量进行直接比较。
*/

/**
interface 由 value 和 type 所表示。
当 value 和 type 同时为 nil 时，interface == nil 才为 true。
（1）【是interface】  在 var err errror;  err = returnErr() ： value 为 nil，但 type 为 *Err，因此与nil比较结果为false;
（2）【不是interface】在 err := returnErr()：err的类型是 *Err，指针为空，与nil比较为true;
（3）【是interface】  在 var err error; err = returnErr(); err == ErrNil：err实际上value 为 nil，type 却为 *Err，本质上是一个Err类型的空指针，与ErrNil比较判断为true

！！！
总结：代码中有使用自定义error类型，需留意err变量在传递过程中，是否赋值到 error interface变量，且比较是否为nil时，要与自定义error的类型进行比较

*/

func TestErrorFormat(t *testing.T) {
	err := errors.New("panic")
	//err = nil
	fmt.Printf("===: %s\n", err)
	fmt.Printf("===: %v\n", err)
	fmt.Printf("===: %+v\n", err)
	fmt.Printf("===: %s\n", err)
	fmt.Printf("===: %s\n", err)
}
