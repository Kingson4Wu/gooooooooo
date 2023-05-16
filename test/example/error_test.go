package example

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

//一文搞懂 Go 错误链
//https://mp.weixin.qq.com/s/nvgNsZgnm_ymb-avI6Zg7Q

func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func processFile(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

func TestFile(t *testing.T) {
	err := processFile("1.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, os.ErrNotExist))
		err = errors.Unwrap(err)
		fmt.Println(err)
		err = errors.Unwrap(err)
		fmt.Println(err)
		return
	}
}

//目前Go标准库中提供的用于wrap error的API有fmt.Errorf和errors.Join。fmt.Errorf最常用，在上面的示例中我们演示过了。errors.Join用于将一组errors wrap为一个error。

func TestFile2(t *testing.T) {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")

	err := fmt.Errorf("wrap multiple error: %w, %w, %w", err1, err2, err3)
	fmt.Println(err)
	e, ok := err.(interface{ Unwrap() []error })
	if !ok {
		fmt.Println("not imple Unwrap []error")
		return
	}
	fmt.Println(e.Unwrap())
}

func TestFile3(t *testing.T) {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")

	err := errors.Join(err1, err2, err3)
	fmt.Println(err)
	errs, ok := err.(interface{ Unwrap() []error })
	if !ok {
		fmt.Println("not imple Unwrap []error")
		return
	}
	fmt.Println(errs.Unwrap())
}

func rootCause(err error) error {
	for {
		e, ok := err.(interface{ Unwrap() error })
		if !ok {
			return err
		}
		err = e.Unwrap()
		if err == nil {
			return nil
		}
	}
}

func TestFile4(t *testing.T) {
	err1 := errors.New("error1")

	err2 := fmt.Errorf("2nd err: %w", err1)
	err3 := fmt.Errorf("3rd err: %w", err2)

	fmt.Println(err3) // 3rd err: 2nd err: error1

	fmt.Println(rootCause(err1)) // error1
	fmt.Println(rootCause(err2)) // error1
	fmt.Println(rootCause(err3)) // error1
}

type MyError struct {
	err string
}

func (e *MyError) Error() string {
	return e.err
}

// 我们通过errors.As将错误链err3中的err1提取到e中，后续就可以使用err1这个特定错误的信息了。
func TestFile5(t *testing.T) {
	err1 := &MyError{"temp error"}
	err2 := fmt.Errorf("2nd err: %w", err1)
	err3 := fmt.Errorf("3rd err: %w", err2)

	fmt.Println(err3)

	var e *MyError
	ok := errors.As(err3, &e)
	if ok {
		fmt.Println(e)
		return
	}
}
