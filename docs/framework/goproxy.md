+ https://github.com/panjf2000/goproxy

```go

func main() {
	goproxy := handler.NewProxyServer()

	log.Infof("Start the proxy server in port:%s", config.RuntimeViper.GetString("server.port"))
	log.Fatal(goproxy.ListenAndServe())
}
```

+ `ln, err := net.Listen("tcp", addr)`
