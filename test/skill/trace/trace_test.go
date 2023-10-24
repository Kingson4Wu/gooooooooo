package trace_test

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

func TestXxx(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Hello, World!")
}

//go tool trace trace.out
