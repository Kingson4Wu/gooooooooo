package synccond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 使用Cond可以避免CPU空转的情况。 !!!!
func TestBroastcast(t *testing.T) {

	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listener goroutines
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}

	go f(10)
	go f(15)

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}

func TestBroastcast2(t *testing.T) {

	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	donation.cond.L.Lock()
	fmt.Println(1)

	go func() {
		time.Sleep(time.Second)
		donation.cond.L.Lock()

		fmt.Println(33)
	}()

	donation.cond.Wait()
	// wait了之后其他goroutine可以lock！！！！

	donation.cond.L.Lock()
	fmt.Println(22)
	donation.cond.L.Lock()
	fmt.Println(2)

}

/**
1、基于channel生产者消费者模式异步操作数据
2、使用Cond避免CPU空转的情况
3、基于消费情况动态调整goroutine的数量
*/
