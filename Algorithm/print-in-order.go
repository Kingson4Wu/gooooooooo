package algorithm

/**

多 Goroutine 的并发程序如何保证按序输出？channel 的使用是关键
http://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651438510&idx=2&sn=cbd1481bab2dfa1839e3ed852634ac95&chksm=80bb635cb7ccea4a00b314499634175a644f57f45949ae471b8a2879a9f24e7a35a4a5612d6a&mpshare=1&scene=1&srcid=&sharer_sharetime=1581472087275&sharer_shareid=dcfe0eae58d1da3d4cc1d60a98c3905c#rd

*/

import (
	"fmt"
)

func First(streamSync [3]chan interface{}) {
	fmt.Print("First")
	streamSync[0] <- nil
}

func Second(streamSync [3]chan interface{}) {
	<-streamSync[0]
	fmt.Print("Second")
	streamSync[1] <- nil
}

func Third(streamSync [3]chan interface{}) {
	<-streamSync[1]
	fmt.Print("Third")
	streamSync[2] <- nil
}

func PrintInOrder(callOrder [3]int) {
	inputCallOrder := callOrder
	fmt.Println("[]inputCallOrder:", inputCallOrder)

	// make an array of unbuffered
	var streamSync [3]chan interface{}
	for i := range streamSync {
		streamSync[i] = make(chan interface{})
	}

	// 建立 [int:func] 對應表
	var functionNumTable = map[int]func([3]chan interface{}){
		1: First,
		2: Second,
		3: Third,
	}

	//依照輸入順序呼叫 goroutine
	for _, fNum := range inputCallOrder {
		go functionNumTable[fNum](streamSync)
	}

	<-streamSync[2]
}

/**
func main() {
	var testCases = [][3]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, theCase := range testCases {
		PrintInOrder(theCase)
		fmt.Println()
		fmt.Println()
	}
}
*/
