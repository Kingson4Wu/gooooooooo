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


+ golang生成wasm,js调用wasm: <https://zhuanlan.zhihu.com/p/618232969>

package main

import (
	"syscall/js"
)

// 导出add函数
func add(this js.Value, inputs []js.Value) interface{} {
	a := inputs[0].Int()
	b := inputs[1].Int()
	return a + b
}

// 导出add函数
func mul(this js.Value, inputs []js.Value) interface{} {
	a := inputs[0].Int()
	b := inputs[1].Int()
	return a * b
}

func main() {
	c := make(chan struct{}, 0)

	// 注册add函数
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("mul", js.FuncOf(mul))

	<-c
}
使用下面的命令编译并打包成wasm。注意golang需要版本>11

set GOOS=js   
set GOARCH=wasm  
go build -o main.wasm main.go
然后就是关键的需要在html中引入的东西了

<meta http-equiv="Content-Security-Policy" content="script-src 'self' 'unsafe-inline' 'unsafe-eval'">
<script src="wasm_exec.js"></script>
<script src="index.js"></script>
需要注意的是,这里的wasm_exec.js最好去你的本地直接拷贝过来

D:\Go\misc\wasm地址下拷贝对应版本的wasm_exec.js到项目中才可以正常使用

然后就是index.js中的调用了

const go = new Go();

WebAssembly.instantiateStreaming(fetch("app/main.wasm"), go.importObject)
    .then((result) => {
        go.run(result.instance);

        // 调用导出函数
        const sum = add(2, 3);
        console.log(`The sum of 2 and 3 is ${sum}.`);

        const ji = mul(2, 3);
        console.log(`The mul of 2 and 3 is ${ji}.`);
    })
    .catch((err) => {
        console.error(err);
    });

// 导入参数对象
const imports = {
    env: {
        log: function (str) {
            console.log(str);
        },
    },
};

// 注册导出函数
function add(x, y) {
    return window.wasmAdd(x, y);
}

// 注册导出函数
function mul(x, y) {
    return window.wasmMul(x, y);
}


----


