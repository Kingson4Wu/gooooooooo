想在mac基于linux开发，可以参考这个？：https://mp.weixin.qq.com/s/WoXZ_i7lJeXny8lIJ_Aj4g

1. goland设置GOOS是linux
2. 使用docker跑服务

为了方便代码跳转到linux的实现，可以修改Goland上的GOOS选项

<pre>
为了能够在Linux下运行代码，可以在机器上安装docker，在容器里跑。docker官方提供了 golang:1.19 镜像，GOPATH 是 /go，我们直接用这个，并把本机的目录映射进去:

# 删除之前的容器，如果有
docker rm -f go_app

# 启动容器
docker run -d \
  --mount type=bind,source=$HOME/go/src,target=/go/src \
  --workdir /go/src/github.com/ \
  --name go_app \
  --restart always \
  golang:1.19 \
  sleep infinity

# 进入容器的命令行
docker exec -it go_app bash

# cd echo_server directory
go run main.go
Golang镜像不提供 netstat vim 等命令，需要手动在容器里安装 net-tools 和 vim:

# 查看linux发行版
cat /etc/issue

# 替换成阿里云 debian 11 的源
# https://developer.aliyun.com/mirror/debian/
cat > /etc/apt/sources.list << EOF \
deb https://mirrors.aliyun.com/debian/ bullseye main non-free contrib \
deb-src https://mirrors.aliyun.com/debian/ bullseye main non-free contrib \
deb https://mirrors.aliyun.com/debian-security/ bullseye-security main \
deb-src https://mirrors.aliyun.com/debian-security/ bullseye-security main \
deb https://mirrors.aliyun.com/debian/ bullseye-updates main non-free contrib \
deb-src https://mirrors.aliyun.com/debian/ bullseye-updates main non-free contrib \
deb https://mirrors.aliyun.com/debian/ bullseye-backports main non-free contrib \
deb-src https://mirrors.aliyun.com/debian/ bullseye-backports main non-free contrib \
EOF

apt update
apt install -y net-tools vim man
</pre>


