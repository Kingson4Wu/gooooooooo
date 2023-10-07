package selectchannl

import (
	"fmt"
	"testing"
	"time"
)

/**
9.4 #64: Expecting deterministic behavior using select and channels（错误的以为 select 的执行结果是确定的）
select 中的 case 执行时机是随机的

解决方案可以是使用for-select嵌套：
*/

func TestR(t *testing.T) {

	messageCh := make(chan int, 10)
	disconnectCh := make(chan struct{})

	//go listing1(messageCh, disconnectCh)
	go listing2(messageCh, disconnectCh)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}
	time.Sleep(5000 * time.Millisecond)

}

func listing1(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

func listing2(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnection, return")
					return
				}
			}
		}
	}
}
