## rune
+ rune类型就是int的别名;是用来区分字符值和整数值的
+ type rune = int32
+ rune属于整型


<pre>
统计带中文字符串长度
// 使用内置函数 len() 统计字符串长度
fmt.Println(len("Go语言编程"))  // 输出：14  
前面说到，字符串在底层的表示是一个字节序列。其中，英文字符占用 1 字节，中文字符占用 3 字节，所以得到的长度 14 显然是底层占用字节长度，而不是字符串长度，这时，便需要用到 rune 类型。

// 转换成 rune 数组后统计字符串长度
fmt.Println(len([]rune("Go语言编程")))  // 输出：6
这回对了。很容易，我们解锁了 rune 类型的第一个功能，即统计字符串长度。

截取带中文字符串
s := "Go语言编程"
// 转成 rune 数组，需要几个字符，取几个字符
fmt.Println(string([]rune(s)[:4])) // 输出：Go语言 

// 统计字符串长度
fmt.Println(utf8.RuneCountInString("Go语言编程")) // 输出：6

</pre>

+ Go 语言把字符分 byte 和 rune 两种类型处理。byte 是类型 unit8 的别名，用于存放占 1 字节的 ASCII 字符，如英文字符，返回的是字符原始字节。rune 是类型 int32 的别名，用于存放多字节字符，如占 3 字节的中文字符，返回的是字符 Unicode 码点值。

+ 在我看来，rune 类型只是一种名称叫法，表示用来处理长度大于 1 字节（ 8 位）、不超过 4 字节（ 32 位）的字符类型。但万变不离其宗，我们使用函数时，无论传入参数的是原始字符串还是 rune，最终都是对字节进行处理。

+ 详解 Go 中的 rune 类型:<https://mp.weixin.qq.com/s/hcrq5fYaQ7FN_2oSMRNjcA>

## 数组
+ 在Go语言中，append导致切片扩容时，如果切片容量小于1000，总是会成倍的扩容底层数组。大于1000时，则会以1.25倍的数量扩容底层数组。

## 函数
+ 带返回值的函数，还可以给返回值命名
```go
func change(a, b int) (x, y int) {
    x = a + 100
    y = b + 100
    return   //101, 102
    //return x, y  //同上
    //return y, x  //102, 101
}
 
```

+ 创建一个匿名函数立即执行
```go
func(nums ...int) {
  fmt.Println(reflect.TypeOf(nums))  // 输出：[]int
  fmt.Println(nums, " ")  // 输出：[1 2 3]
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total) // 输出：6
}(1, 2, 3)

```


----


+ Go语言进阶之路（二）：字符串和指针:<https://blog.csdn.net/c315838651/article/details/105008272>

---

+ golang 函数作为参数传递(回调):<https://blog.csdn.net/MaxCoderLlj/article/details/80191998>

```go
func sayhello(name string, f func(string, string)) {
    f("hello", name)
}
```