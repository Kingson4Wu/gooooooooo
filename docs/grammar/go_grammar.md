
+ /gooooooooo/test/edu/gobyexample !!!!!

----

+ https://www.runoob.com/go/go-constants.html

+ package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
+ 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）
+ Go 代码中会使用到的 25 个关键字或保留字;还有 36 个预定义标识符
+ 变量默认值
<pre>
数值类型（包括complex64/128）为 0
布尔类型为 false
字符串为 ""（空字符串）
其他几种类型为 nil
</pre>
+ 根据值自行判定变量类型
+ := 是一个声明语句.这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
+ 多变量声明
<pre>
var vname1, vname2, vname3 = v1, v2, v3
vname1, vname2, vname3 := v1, v2, v3'
// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
</pre>
+ 值类型和引用类型
    - 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
    - 你可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中。内存地址为称之为指针，这个指针实际上也被存在另外的某一个字中.
+ 全局变量是允许声明但不使用。 同一类型的多个变量可以声明在同一行;多变量可以在同一行进行赋值(并行 或 同时 赋值)
+ 如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
+ 空白标识符 _ 
<pre>
空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。
</pre>

---

+ 常量
    - const b = "abc"
    - const c_name1, c_name2 = value1, value2
+ 常量还可以用作枚举;常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值,常量表达式中，函数必须是内置函数
<pre>
const (
    Unknown = 0
    Female = 1
    Male = 2
)
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)
</pre>    
+ iota - iota，特殊常量，可以认为是一个可以被编译器修改的常量。
<pre>
const (
    a = iota
    b = iota
    c = iota
)
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式

const (
    a = iota
    b
    c
)

const (
    i=1<<iota
    j=3<<iota
    k
    l
)
i= 1,j= 6,k= 12,l= 24
</pre>

+ Go 语言运算符 ：跟java基本相同
    - & 返回变量存储地址
    - * 指针变量。

+ Go 语言条件语句
    - if...else 语句
    - switch 语句
    - select 语句:select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
    - Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

+ Go 语言循环语句
    - for 循环,break, continue, goto
    - 无限循环 for true  { ... }

+ Go 语言函数, 函数作为另外一个函数的实参
<pre>
func function_name( [parameter list] ) [return_types] {
   函数体
}
函数返回多个值
func swap(x, y string) (string, string) {
   return y, x
}
</pre>    

+ Go 语言变量作用域(没啥特别)
+ Go 语言数组 
<pre>
var balance [10] float32
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
</pre>

+ Go 语言指针
    - 变量是一种使用方便的占位符，用于引用计算机内存地址。Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
    - Go 空指针: 当一个指针被定义后没有分配到任何变量时，它的值为 nil;nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
    - 空指针判断 : if(ptr != nil) 

+ Go 语言结构体
<pre>
type Books struct {
   title string
   author string
   subject string
   book_id int
}
结构体指针 *Books
</pre>

构造函数
<pre>
//MinStack is
type MinStack struct {
	array  []int
	minVal int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{
		array:  []int{},
		minVal: 0,
	}
}
</pre>

---
+ Go 语言切片(Slice)
    - Go 语言切片是对数组的抽象。
    - len() 和 cap() 函数,cap() 可以测量切片最长可以达到多少
    - append() 和 copy() 函数:如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
    - 空(nil)切片: 一个切片在未初始化之前默认为 nil，长度为 0
<pre>
容量capacity,len 是数组的长度并且也是切片的初始长度
s :=make([]int,len,cap) 
var numbers = make([]int,3,5)
子切片从索引1(包含) 到索引4(不包含)
numbers[1:4]
/* 向切片添加一个元素 */
   numbers = append(numbers, 1)
/* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
/* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
/* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)      
</pre>    

<pre>
在数组头前插入元素
		result = append(result, nil)
		copy(result[1:], result[0:])
		result[0] = column
</pre>

+ Go 语言范围(Range) :Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
<pre>
//这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时候我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
</pre>    

+ Go 语言Map(集合)
<pre>
countryCapitalMap = make(map[string]string)
countryCapitalMap [ "France" ] = "巴黎"
/*使用键输出地图值 */
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }
 /*查看元素在集合中是否存在 */
    capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("American 的首都是", capital)
    } else {
        fmt.Println("American 的首都不存在")
    }    
//delete() 函数
delete(countryCapitalMap, "France")

</pre>

+ Go 语言递归函数
+ Go 语言类型转换
    -  var sum int = 17 ; float32(sum)
+ Go 语言接口 
<pre>
type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()
}
</pre>   

+ Go 错误处理
    - Go 语言通过内置的错误接口提供了非常简单的错误处理机制。error类型是一个接口类型
    - 我们可以在编码中通过实现 error 接口类型来生成错误信息。
        1. 函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息
        2. 通过实现实现 `error` 接口接口
    - <https://www.runoob.com/go/go-error-handling.html>

+ Go 并发
    - Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的        
 <pre>
 func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

func main() {
        go say("world")
        say("hello")
}
 </pre>   
+ 通道（channel）:用来传递数据的一个数据结构 !!!!
    - 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
    - 默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须又接收端相应的接收数据。
<pre>
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
//声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建
ch := make(chan int)           
</pre>
+ https://www.runoob.com/go/go-concurrent.html
+ 先进后出？？？！why


---
+ byte is an alias for uint8 and is equivalent to uint8 in all ways    
+ non-constant array bound length go
    - You can't instantiate an array like that with a value calculated at runtime. Instead use make to initialize a slice with the desired length. It would look like this;
    left := make([]int, leftLength)

---
1. 字符数组和字符串互转
sa := []rune(a) 
s := string(sa)   
 
 ---
 + 切片中间插入元素
 ```go
 		rear := append([]int{}, this.deque[index:]...)
		this.deque = append(this.deque[:index], value)
		this.deque = append(this.deque, rear...)
 ```

 ---

### make vs new
+ <http://www.flysnow.org/2017/10/23/go-new-vs-make.html>
+ 二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；而new用于类型的内存分配，并且内存置为零。所以在我们编写程序的时候，就可以根据自己的需要很好的选择了。
+ make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
+ new这个内置函数，可以给我们分配一块内存让我们使用，但是现实的编码中，它是不常用的。我们通常都是采用短语句声明以及结构体的字面量达到我们的目的; make函数是无可替代的，我们在使用slice、map以及channel的时候，还是要使用make进行初始化，然后才才可以对他们进行操作.

### 数组 vs 切片
+ 在 Go 语言中，声明变量类型的方式有好几种，比如使用关键词 var ，又比如使用字面量方式，如果要声明的变量是一个集合类型，则还需要用到 [] 操作符，其中，在 [] 操作符中设定数值的叫数组，如：var [5]int；未设定数值的叫切片，如：var []int。
+ 数组长度设定后不可更改，也由此说它是固定长度序列。
+ 先有数组才有切片，切片是基于数组创建的一种数据结构。结构体的定义在 reflect.SliceHeader 中，它由 3 个字段组成，分别表示指向底层数组的指针、长度、容量。切片跟数组最大的不同是其长度可变，可扩容，可缩减，因此大家也叫它“动态数组”。
+ nil 切片、空切片与零切片是切片的三种状态，nil 切片是指在声明时未做初始化的切片，不用分配内存空间，一般使用 var 创建。使用 make 创建的空切片需要分配内存空间，nil 切片与空切片的长度、容量都为 0 ，如果我们要创建长度容量为 0 的切片，官方推荐 nil 切片。零切片指初始值为类型零值的切片。
+ Go 语言提供了内置函数 copy() 用来拷贝切片。
+ copy() 函数的操作使用 append() 函数也能实现，只不过 append() 函数每次复制数据都需要创建临时切片，相比性能上 copy() 函数更胜一筹。
+ 再谈Go数组与切片：<https://mp.weixin.qq.com/s/4kqaKDDFUCjQ2T9itWlidw>

---

### Go语言圣经（中文版）
+ <https://books.studygolang.com/gopl-zh/>