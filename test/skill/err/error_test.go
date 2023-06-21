package errrr

import (
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
