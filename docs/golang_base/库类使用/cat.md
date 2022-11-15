+ catyjianshu.com/p/96c3232b9d7e

+ 客户端篇
步骤1： 客户端机器创建文件
mkdir /data
chmod -R 777 /data/
mkdir -p /data/appdatas/cat/
cd /data/appdatas/cat/
vim client.xml 
将下列配置复制到client.xml中
注意: 2280是默认的CAT服务端接受数据的端口，不允许修改，http-port是Tomcat启动的端口，默认是8080，建议使用默认端口


<?xml version="1.0" encoding="utf-8"?>
<config mode="client">
    <servers>
        <server ip="10.1.1.1" port="2280" http-port="8080"/>
        <server ip="10.1.1.2" port="2280" http-port="8080"/>
        <server ip="10.1.1.3" port="2280" http-port="8080"/>
    </servers>
</config>

作者：枫之鬼影
链接：https://www.jianshu.com/p/96c3232b9d7e
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

+ https://github.com/dianping/cat
+ https://github.com/dianping/cat/blob/master/lib/go/README.zh-CN.md
