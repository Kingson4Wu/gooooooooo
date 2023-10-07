package slice

import (
	"fmt"
	"runtime"
	"testing"
)

func TestAppend(t *testing.T) {

	listing1()
	fmt.Println("-----")
	listing2()
	fmt.Println("-----")
	listing3()
}

// 当调用f函数的时候，对s入参做的修改都会反应到上层函数中的s
func listing1() {
	s := []int{1, 2, 3}

	f(s[:2])
	fmt.Println(s)
}

// 解决方案一：对s进行深拷贝
func listing2() {
	s := []int{1, 2, 3}
	sCopy := make([]int, 2)
	copy(sCopy, s)

	f(sCopy)
	result := append(sCopy, s[2])
	fmt.Println(result)
}

// 解决方案二：使用 full slice expression: s[low:high:max]，这样就限制了切片的cap
// cap = max-low = 2-0 = 2
// 这样f函数对切片进行append操作就不用影响其他共享底层存储的切片了
func listing3() {
	s := []int{1, 2, 3}
	f(s[:2:2])
	fmt.Println(s)
}

func f(s []int) {
	_ = append(s, 10)
}

/**

来自chatgpt：

在Go语言中，full slice expression（完整切片表达式）是一种用于创建切片的表达式，
它可以提取底层数组中的一段元素序列，并创建一个新的切片。这种表达式的语法如下：

```go
slice[low:high:max]
```

- `slice` 是要切片的数组或切片。
- `low` 是切片的起始索引，它指定了切片的第一个元素。
- `high` 是切片的结束索引，它指定了切片的最后一个元素（不包括这个索引对应的元素）。
- `max` 是切片的容量，它限制了切片的最大长度，通常是底层数组中从 `low` 开始的元素数量。

使用完整切片表达式时，创建的新切片将包含底层数组中从 `low` 索引开始到 `high-1` 索引结束的元素。新切片的长度为 `high - low`，容量为 `max - low`。

以下是一个示例，演示了如何使用完整切片表达式：

```go
package main

import "fmt"

func main() {
    // 创建一个切片
    original := []int{1, 2, 3, 4, 5}

    // 使用完整切片表达式创建一个新切片
    newSlice := original[1:3:4] // 从索引1开始，到索引3-1=2结束，容量为4-1=3

    fmt.Println(newSlice)       // 输出: [2 3]
    fmt.Println(len(newSlice))   // 输出: 2 (长度为2)
    fmt.Println(cap(newSlice))   // 输出: 3 (容量为3)
}
```

在上述示例中，`original[1:3:4]` 创建了一个新的切片，
它包含了原始切片 `original` 中从索引1到索引2的元素（不包括索引2），
并且它的容量被限制在3，因此新切片的长度为2，容量为3。
这种灵活的切片操作可以用于优化内存使用或创建切片的不同视图。
*/

//------------------------

func TestLeak(t *testing.T) {
	foos := make([]Foo, 1_000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	// 和上面类似不会回收剩下998个元素
	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)

	printAlloc()

	fmt.Println("=====")

	//t2 := keepFirstTwoElementsOnlyCopy(foos)
	t2 := keepFirstTwoElementsOnlyMarkNil(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(t2)

}

/**
141 KB
1024167 KB
1024168 KB
1024169 KB
=====
2193 KB
*/

type Foo struct {
	v interface{}
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	return foos[:2]
}

// keepFirstTwoElementsOnlyCopy 和 keepFirstTwoElementsOnlyMarkNil
// 都能保证剩下的资源被回收，但是如果想留下的元素不是2，我们想保留500个元素，
// 使用 keepFirstTwoElementsOnlyMarkNil 会有更高的效率，不用额外拷贝资源
func keepFirstTwoElementsOnlyCopy(foos []Foo) []Foo {
	res := make([]Foo, 2)
	copy(res, foos)
	return res
}

// keepFirstTwoElementsOnlyMarkNil 给元素中的指针成员置为nil，能够让GC自动回收
func keepFirstTwoElementsOnlyMarkNil(foos []Foo) []Foo {
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
