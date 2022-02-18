package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

/**
这里目前是一个简易版本的Socket5 代理，还有很多没有实现，要实现完成的Socket5，可以自己根据他定义的协议试试。
*/
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%d\n", n)

	if b[0] == 0x05 { //只处理Socket5协议
		//客户端回应：Socket服务端不需要验证方式
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		if err != nil {
			log.Println(err)
			return
		}
		var host, port string
		switch b[3] {
		case 0x01: //IP V4
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03: //域名
			host = string(b[5 : n-2]) //b[4]表示域名的长度
		case 0x04: //IP V6
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		}
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))

		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			log.Println(err)
			return
		}
		defer server.Close()
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端连接成功
		//进行转发
		go io.Copy(server, client)
		io.Copy(client, server)
	}

}

/**
http://www.flysnow.org/2016/12/26/golang-socket5-proxy.html
Golang实现一个Socket5的简单代理。Socket5和HTTP并没有太大的不同，他们都可以完全给予TCP协议，只是请求的信息结构不同，所以这次我们不能像上次HTTP Proxy一样，解析请求和应答，要按照Socket的协议方式解析。
Socket协议分为Socket4和Socket5两个版本，他们最明显的区别是Socket5同时支持TCP和UDP两个协议，而SOcket4只支持TCP。目前大部分使用的是Socket5
*/
