package slice

import (
	"fmt"
	"testing"
)

//Go Slice 扩容的这些坑你踩过吗 : https://mp.weixin.qq.com/s/fbToVV_aC2rYk_C4S3doZA

/**
切片底层结构定义：包含指向底层数组的指针、长度和容量

type slice struct {
  array unsafe.Pointer
  len   int
  cap   int
}
append 操作：可以是 1 个、多个、甚至整个切片（记得后面加...）；添加元素时当容量不足，则会自动触发切片扩容机制，产生切片副本，同时指向底层数组的指针发生变化
*/

// 案例 1：传值+未扩容
func TestXxxx1(t *testing.T) {
	s1 := make([]int, 0, 5)
	fmt.Printf("s1切片地址: %p\n", s1)
	fmt.Println("s1切片: ", s1)
	appendFunc(s1)
	fmt.Println("s1切片: ", s1)
	fmt.Println("s1切片表达式: ", s1[:5])
}

func appendFunc(s2 []int) {
	s2 = append(s2, 1, 2, 3)
	fmt.Printf("s2切片地址: %p\n", s2)
	fmt.Println("s2切片: ", s2)

}

/**
s1切片:  []
s2切片:  [1 2 3]
s1切片:  []
s1切片表达式:  [1 2 3 0 0]
*/

/**
通过分析 fmt 包的源码，不难发现，打印的地址，其实是切片里指向底层数组的指针存储的地址，并不是两个切片本身的地址。同时也说明这两切片是指向同一个底层数组。


原因正式分析：


传值操作，s1 和 s2 是两个不同的切片变量，但是指向底层数组的指针是同一个；

长度和容量的变化：s1 Len=0 和 Cap=5，后来未发生过变化；s2 一开始被赋值时 Len=0 和 Cap=5，在 append 操作后，Len=3 和 Cap=5，同时底层数组值从[0,0,0,0,0]被修改成了[1,2,3,0,0];

输出结果，s1 由于 Len=0 所以输出空[]，而 s1 用切片表达式，是基于底层数组[1,2,3,0,0]进行切片，所以输出结果为[1,2,3,0,0]；
*/

// 案例 2：传值+扩容
func TestXxxx2(t *testing.T) {
	s1 := make([]int, 0, 5)
	fmt.Println("s1切片: ", s1)
	appendFunc2(s1)
	fmt.Println("s1切片: ", s1)
	fmt.Println("s1切片表达式: ", s1[:5])
}

func appendFunc2(s2 []int) {
	s2 = append(s2, 1, 2, 3, 4, 5, 6)
	fmt.Println("s2切片: ", s2)
}

/**
s1切片:  []
s2切片:  [1 2 3 4 5 6]
s1切片:  []
s1切片表达式:  [0 0 0 0 0]
*/

func TestXxxx2_(t *testing.T) {
	s1 := make([]int, 0, 5)
	fmt.Println("s1切片: ", s1)
	appendFunc2_(s1)
	fmt.Println("s1切片: ", s1)
	fmt.Println("s1切片表达式: ", s1[:5])
}

func appendFunc2_(s2 []int) {
	s2 = append(s2, 1, 2, 3, 4, 5)
	s2 = append(s2, 6)
	fmt.Println("s2切片: ", s2)
}

/**
s1切片:  []
s2切片:  [1 2 3 4 5 6]
s1切片:  []
s1切片表达式:  [1 2 3 4 5]
*/

/**
原因分析：

发生扩容后，s2 指向的底层数组会产生副本，导致 s1 和 s2 不再指向同一个底层数组；

长度和容量的变化：s2 append 后 Len=6、Cap=10 和底层数组值为[1,2,3,4,5,6,0,0,0,0]；s2 的操作完全不影响 s1 的数据，s1 仍然是 Len=0、Cap=5 和底层数组值为[0,0,0,0,0]；

输出结果，s2 由于 Len=6 所以输出[1,2,3,4,5,6]，s1 由于 Len=0 所以输出空[]，而 s1 用切片表达式，是基于底层数组[0,0,0,0,0]进行切片，所以输出结果为[0,0,0,0,0]；
*/

//案例 3：传址+不关心扩容

//上面两个传值操作的例子，不管扩容与否，都不会影响原切片 s1 的长度和容量。如果我们期望修改 s2 的同时也修改原切片 s1，则需要用到切片指针，基于地址传递进行操作。

func TestXxxx3(t *testing.T) {
	s1 := make([]int, 0, 5)
	fmt.Println("s1切片: ", s1)
	fmt.Printf("s1切片地址: %p len:%d cap:%d\n", &s1, len(s1), cap(s1))
	appendFunc3(&s1)
	fmt.Println("s1切片: ", s1)
	fmt.Println("s1切片表达式: ", s1[:5])
}

func appendFunc3(s2 *[]int) {
	fmt.Printf("s2切片地址: %p len:%d cap:%d\n", s2, len(*s2), cap(*s2))
	//*s2 = append(*s2, 1, 2, 3)
	*s2 = append(*s2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("append后s2切片地址: %p len:%d cap:%d\n", s2, len(*s2), cap(*s2))
	fmt.Println("s2切片: ", *s2)
}

/**
s1切片:  []
s1切片地址: 0x1400009a0c0 len:0 cap:5
s2切片地址: 0x1400009a0c0 len:0 cap:5
append后s2切片地址: 0x1400009a0c0 len:10 cap:10
s2切片:  [1 2 3 4 5 6 7 8 9 10]
s1切片:  [1 2 3 4 5 6 7 8 9 10]
s1切片表达式:  [1 2 3 4 5]
*/

//万变不离其宗，传址操作，始终操作的是同一个切片变量，append 操作后，长度和容量都会同时发生变化，以及如果触发扩容，那么指向底层数组的指针，也都会同时发生变化。

/**
总结

切片传值操作，append 未触发扩容，会同时修改底层数组的值，但不会影响原切片的长度和容量；当触发扩容，那么会产生副本，后面的修改则会和原底层数组剥离开，互不影响。

如果期望在修改切片后，对原切片也发生修改，则可以使用传址操作，始终基于同一个切片变量进行操作。
*/
