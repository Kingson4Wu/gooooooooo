https://www.cnblogs.com/niuben/p/14461973.html

为什么需要 embed 包
在以前，很多从其他语言转过来Go语言的同学会问到，或者踩到一个坑。就是以为Go语言所打包的二进制文件中会包含配置文件的联同编译和打包。

结果往往一把二进制文件挪来挪去，就无法把应用程序运行起来了。因为无法读取到静态文件的资源。

无法将静态资源编译打包二进制文件的话，通常会有两种解决方法：

第一种是识别这类静态资源，是否需要跟着程序走。
第二种就是将其打包进二进制文件中。
第二种情况的话，Go以前是不支持的，大家就会借助各种花式的开源库，例如：go-bindata/go-bindata来实现。

但是在Go1.16起，Go语言自身正式支持了该项特性。

它有以下优点

能够将静态资源打包到二进制包中，部署过程更简单。传统部署要么需要将静态资源与已编译程序打包在一起上传，或者使用docker和dockerfile自动化前者，这是很麻烦的。
确保程序的完整性。在运行过程中损坏或丢失静态资源通常会影响程序的正常运行。
静态资源访问没有io操作，速度会非常快。


在embed中，可以将静态资源文件嵌入到三种类型的变量，分别为：字符串、字节数组、embed.FS文件类型

//go:embed version.txt
var version string

//go:embed version.txt
var versionByte []byte

//go:embed static
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
    if useOS {
        log.Print("using live mode")
        return http.FS(os.DirFS("static"))
    }

    log.Print("using embed mode")

    fsys, err := fs.Sub(embededFiles, "static")
    if err != nil {
        panic(err)
    }
    return http.FS(fsys)
}

