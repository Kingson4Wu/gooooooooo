package copy

import (
	"fmt"
	"testing"
)

/**

3.8 #24: Not making slice copies correctly（没有正确的复制切片）
copy函数会复制元素的个数为2个切片长度最小值

[]
-------
[0 1 2]
*/

func TestCopy(t *testing.T) {

	bad()
	fmt.Println("-------")
	correct()
}

func bad() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst)

	_ = src
	_ = dst
}

func correct() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst)

	_ = src
	_ = dst
}
