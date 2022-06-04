+  json.Marshal
Marshal函数，把Go对象转换成JSON字节切片
+ json.Unmarshal
Unmarshal函数，把JSON字节切片转换成Go对象（不可导出的变量无法解析）
+ json.NewEncoder&Encoder.Encode
json.Encoder，将Encode方法参数的Go结构体对象，编码成JSON内容输出到Writer中。
+ json.NewDecoder&Decoder.Decode
Decoder，将Reader中的JSON内容解码成Go对象，存入Decode方法的参数中，Decode方法的参数一般传入对象的地址（指针类型）。