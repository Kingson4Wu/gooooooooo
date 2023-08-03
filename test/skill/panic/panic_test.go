package panic_test

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {

	//f()
	f2()
	fmt.Println("success ... ")

	//!!!!  新开goroutine的panic无法recover  ！！！

	ch := make(chan bool)
	<-ch
}

func f() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if true {
		panic("xxx")
	}
}

func f2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("====")
			fmt.Println(err)
		}
	}()

	if true {
		go func() {
			panic("oooo")
		}()
	}
	fmt.Println("==000==")
}
