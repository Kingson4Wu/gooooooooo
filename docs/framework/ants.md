
+ golang ants协程池源码分析:<https://www.jianshu.com/p/440f4c3f7c78>

---


+ example.go

```go

ackage main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	fmt.Printf("free goroutines: %d\n", ants.Free())
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
	/** 跑完了为什么还是1000呢？？！！！ */
	/** 表示可用的goroutine数目，但现在可能是空闲的 */
	fmt.Printf("free goroutines: %d\n", ants.Free())

	// Use the pool with a method,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	fmt.Printf("free goroutines: %d\n", p.Free())
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	fmt.Printf("free goroutines=== %d\n", p.Free())
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

	/** 999 * 1000 / 2 */
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}
}

```

+ `runtime.GOMAXPROCS(0)` 获取go程序能使用的cpu核数

+ 全局变量在被import的时候已经被初始化！！
Init an instance pool when importing ants.
	defaultAntsPool, _ = NewPool(DefaultAntsPoolSize)

+ Gosched
runtime.Gosched() 用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

+ 自旋锁

+ ants的实现：

```go
package internal

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

const maxBackoff = 16

func (sl *spinLock) Lock() {
	backoff := 1
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		// Leverage the exponential backoff algorithm, see https://en.wikipedia.org/wiki/Exponential_backoff.
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock instantiates a spin-lock.
func NewSpinLock() sync.Locker {
	return new(spinLock)
}

//自旋锁5次之后自动获得锁！
```    

+ goroutine阻塞等待任务，达到复用效果 ！！！！
本质上还是通过chan来阻塞goroutine达到复用目的

```go

      // 阻塞等待task， - w.task -> task 
      for f := range w.task {
         // task被过期清空或者释放
         if f == nil {
			 //结束goroutine
            return
         }
		 //执行任务！！！
         f()
         // 将w存到pool.workers中下次可以再次获取
         if ok := w.pool.revertWorker(w); !ok {
			 //结束goroutine
            return
         }
      }
```