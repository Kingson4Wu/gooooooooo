package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

		//Avoid re-use of the same i value in each goroutine closure. See the FAQ for more details.
		//https://golang.org/doc/faq#closures_and_goroutines
        i := i

        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    wg.Wait()

}

//This WaitGroup is used to wait for all the goroutines launched here to finish. Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.

//Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.

//Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.

//Note that this approach has no straightforward way to propagate errors from workers. For more advanced use cases, consider using the errgroup package.
//https://pkg.go.dev/golang.org/x/sync/errgroup

/**

To bind the current value of v to each closure as it is launched, one must modify the inner loop to create a new variable each iteration. One way is to pass the variable as an argument to the closure:

    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
In this example, the value of v is passed as an argument to the anonymous function. That value is then accessible inside the function as the variable u.

Even easier is just to create a new variable, using a declaration style that may seem odd but works fine in Go:

    for _, v := range values {
        v := v // create a new 'v'.
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }
This behavior of the language, not defining a new variable for each iteration, may have been a mistake in retrospect. It may be addressed in a later version but, for compatibility, cannot change in Go version 1.

*/