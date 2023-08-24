package main

// #include "atomic_add.h"
import "C"
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var value int64

	numIterations := 100000
	wg.Add(numIterations)

	for i := 0; i < numIterations; i++ {
		go func() {
			C.AtomicAdd((*C.longlong)(&value), 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", value)
}
