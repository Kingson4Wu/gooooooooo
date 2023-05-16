+ 如何在 Go 中使用 Wasm：浅聊 WebAssembly:<https://juejin.cn/post/7195217413091262523>


### gotip

+ https://pkg.go.dev/golang.org/dl/gotip

+ go install golang.org/dl/gotip@latest
+ gotip download
+ GOARCH=wasm GOOS=wasip1 gotip build -o main.wasm main.go

开源的wasm运行时有很多，wazero[17]是目前比较火的且使用纯Go实现的wasm运行时程序，安装wazero后，可以用来执行上面编译出来的main.wasm：

curl -O  https://wazero.io/install.sh
sh install.sh
./bin/wazero run main.wasm
hello

----

