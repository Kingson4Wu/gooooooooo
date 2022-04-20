## 文件
### Go语言获取路径的文件名、后缀
```go
files := "E:\\data\\test.txt"
    paths, fileName := filepath.Split(files)
    fmt.Println(paths, fileName)      //获取路径中的目录及文件名 E:\data\  test.txt
    fmt.Println(filepath.Base(files)) //获取路径中的文件名test.txt
    fmt.Println(path.Ext(files))      //获取路径中的文件的后缀 .txt
```

### 读文件全部内容：
+ `ioutil.ReadAll(file)`

### 创建文件
+ `ofile, _ := os.Create(oname)`
    - 按格式写文件：`fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)`

## 数组
### 在数组头前插入元素
```go
		result = append(result, nil)
		copy(result[1:], result[0:])
		result[0] = column
```

## 切片
### 切片中间插入元素
 ```go
 		rear := append([]int{}, this.deque[index:]...)
		this.deque = append(this.deque[:index], value)
		this.deque = append(this.deque, rear...)
 ```

## 字符串
### 字符数组和字符串互转
```go
sa := []rune(a) 
s := string(sa)
```   

### golang中判断字符串是否为空的方法
```go
if len(str) == 0{
}
//或使用下面的方法判断：
if str == "" {
}
```

## 打印
+ Go语言的%d,%p,%v等占位符的使用：<https://www.jianshu.com/p/66aaf908045e>
    - `fmt.Printf("%v,%v,%v,%v\n", os.Args[0], os.Args[1], os.Args[2], os.Args[3])`

## 排序
```go

type KeyValue struct {
	Key   string
	Value string
}

type ByKey []mr.KeyValue

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

intermediate := []mr.KeyValue{}
sort.Sort(ByKey(intermediate))

```

## HTTP
### Go 语言下载文件 http.Get() 和 io.Copy()
+ <https://www.twle.cn/t/384>

## logger
+ <https://blog.csdn.net/xmcy001122/article/details/119916227>

---

+ Go 语言切片是对数组的抽象。make

