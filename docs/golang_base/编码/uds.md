unix domain socket 

https://zhuanlan.zhihu.com/p/426644841


使用域套接字实现的HTTP服务器可以同时接受多个请求。在Go语言中，可以使用标准库中的net和net/http包来创建一个HTTP服务器，并将其绑定到一个域套接字上。

当你使用域套接字作为HTTP服务器的监听地址时，它可以同时处理多个并发连接。每当有新的请求到达服务器时，Go语言会自动创建一个新的goroutine来处理该请求，从而实现并发处理。

下面是一个简单的示例，演示了如何使用域套接字实现一个基本的HTTP服务器：

go
Copy code
package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	// 创建域套接字
	socketPath := "/tmp/mysocket.sock"
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer l.Close()

	// 处理HTTP请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// 启动HTTP服务器
	server := &http.Server{}
	server.Serve(l)
}
在上述示例中，net.Listen("unix", socketPath)用于创建域套接字监听器。然后，我们通过http.HandleFunc注册一个处理HTTP请求的处理函数，并将其绑定到根路径"/"。最后，通过调用server.Serve(l)来启动HTTP服务器并开始监听来自域套接字的请求。

这样实现的HTTP服务器可以同时接受和处理多个请求。每个请求都会在独立的goroutine中处理，以实现并发性。需要注意的是，你可能需要适当地处理错误、优化并发处理和管理连接的生命周期，以构建更健壮和高效的HTTP服务器。


