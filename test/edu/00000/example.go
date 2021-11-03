package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	//c := make(chan int)
	c := make(chan int, 100)

	/** 加了go关键字之后，顺序是相反的，不加却是正常顺序 */
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//go sum([]int{4, 5, 6}, c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
	//x := <-c
	//fmt.Println(x)

	c <- 1
	c <- 2
	a, b := <-c, <-c
	fmt.Println(a, b)

	go fmt.Println(888)
	go fmt.Println(999)
	go fmt.Println(1000)
}
