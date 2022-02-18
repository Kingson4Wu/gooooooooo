<https://www.jianshu.com/p/bfaba9b6d46d>

编译代码命令

env GOOS=linux GOARCH=386 go build main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

PS：这里386是一个很大的坑，这里是要运行这个打包后程序的平台。如果是linux需要你用 uname -a 来查看你运行的linux系统环境。常见的环境一般有 amd64，i386等

作者：IT锟
链接：https://www.jianshu.com/p/bfaba9b6d46d
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


编译完成后会生成一个 main可执行文件，没有后缀，这时你只需要把这个文件上传到你的虚拟机，直接运行就好了

----

生产环境建议大家使用:
4543是要结束的进程pid
singo是二进制文件名
使用&&同时执行2条命令，避免服务中断
kill 4543 && nohup ./singo > nohup.log 2>&1 &

开发环境强制结束进程，可以使用 kill -9 :
4543是要结束的进程pid
singo是二进制文件名
使用&&同时执行2条命令，避免服务中断
kill -9 4543 && nohup ./singo > nohup.log 2>&1 &
