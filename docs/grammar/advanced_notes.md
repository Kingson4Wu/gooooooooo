## rune
+ rune类型就是int的别名;是用来区分字符值和整数值的
+ type rune = int32
+ rune属于整型

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