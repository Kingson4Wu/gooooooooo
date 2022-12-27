+ wget https://github.com/coder/code-server/releases/download/v4.9.1/code-server-4.9.1-arm64.rpm

<pre>
[root@centos labali]# export PASSWORD="password" && /usr/bin/code-server --port 9999 --host 0.0.0.0
node:internal/modules/cjs/loader:1183
  return process.dlopen(module, path.toNamespacedPath(filename));
                 ^

Error: /lib64/libc.so.6: version `GLIBC_2.25' not found (required by /usr/lib/code-server/node_modules/argon2/lib/binding/napi-v3/argon2.node)
    at Object.Module._extensions..node (node:internal/modules/cjs/loader:1183:18)
    at Module.load (node:internal/modules/cjs/loader:981:32)
    at Function.Module._load (node:internal/modules/cjs/loader:822:12)
    at Module.require (node:internal/modules/cjs/loader:1005:19)
    at require (node:internal/modules/cjs/helpers:102:18)
    at Object.<anonymous> (/usr/lib/code-server/node_modules/argon2/argon2.js:6:25)
    at Module._compile (node:internal/modules/cjs/loader:1101:14)
    at Object.Module._extensions..js (node:internal/modules/cjs/loader:1153:10)
    at Module.load (node:internal/modules/cjs/loader:981:32)
    at Function.Module._load (node:internal/modules/cjs/loader:822:12) {
  code: 'ERR_DLOPEN_FAILED'
</pre>

+ centos服务器安装code-server
https://cloud.tencent.com/developer/article/1655175


+ https://github.com/coder/code-server/releases