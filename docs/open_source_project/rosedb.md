+ https://github.com/flower-corp/rosedb


## bitcask
+ bitcask: 简洁且能快速写入的存储系统模型:<https://zhuanlan.zhihu.com/p/551334186>

+ 在软件系统中, 数据通常存储于内存或磁盘,
完全将数据存储于内存的系统, 比如redis,
面向磁盘的存储系统, 比如数据库mysql, 存储ceph等系统,
数据存储于内存的系统速度更快, 适合作为缓存,
面向磁盘的存储系统则可以存储比内存大得多的数据, 适合大容量存储,
面向磁盘的存储系统需要一个存储引擎管理存储于磁盘中的数据,
存储引擎的实现主要基于两种算法, B+树或lsm,
在数据结构及算法相关课程中我们可能了解过B+树,
知道树一类算法在检索上效率比较高,
基于B+树进行搜索, 在读取数据上速度更快,
适合读多写少的系统,
bitcask是一种类似lsm实现的算法
bitcask简化了lsm的实现逻辑, 简化了代码实现
bitcask相对于B+树一类算法也能做到更高地写入效率

+ 内存中数据结构为keydir, 采用哈希表实现,
每个key指向日志文件id,
+ 磁盘文件由一个active file和多个older file组成,
active file为当前写入的活动中的日志文件,
older file为不再修改的日志文件
+ 数据元组是存储在磁盘日志中的基本单元,
+ 如何能快速写入
bitcask的写入非常简单也很高效, 采用直接追加写入日志的方式,
+ 读取操作也很简单, bitcask先在内存中从keydir根据key得到value的file id和value位置,
然后再到对应的磁盘中文件位置读取value值
+ 删除操作也非常快速, 先从内存中keydir删除对应key, 然后在日志中追加一个数据删除的标志, (两个操作的原子性问题????)
+ 如何解决日志数据冗余
bitcask无论是插入数据, 还是删除数据, 都是采用追加日志的方法,
这样虽然很高效, 但是问题是会在日志中有很多冗余数据,
这时就需要一个定期的日志压缩算法, 将冗余的日志数据合并遍历全部的older file, 如果数据元组中的key在keydir存在, 并且其时间戳与keydir中相同,
说明是最新的数据, 写入merged data file, 同时将:时间戳 key长度 value长度 value位置 key写入一个hint file
+ 额外的好处
使用日志记录操作数据的方式, 还有一种额外的好处,
就是当系统崩溃/重启后, 只需读取hint file即可重新构建出keydir,
将数据信息快速恢复出来.

### bitcask的paper


---

+ RoseDB源码阅读:<https://iocing.github.io/2022/08/15/roseDB%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB/>


+ 还有命令提示, 感觉很棒
+ access data via cli(a copy of redis-cli) 
+ 好像是直接用redis的,直接连redis也是可以的

+ https://redis.io/docs/manual/cli/
+ https://download.redis.io/redis-stable.tar.gz
+ redis-7.0.7.tar.gz

+ mac 安装 redis-cli: https://blog.csdn.net/xingjia001/article/details/108622284


### 文件锁
+ flock/flock_unix.go AcquireFileLock

flock 主要三种操作类型：
LOCK_SH：共享锁，多个进程可以使用同一把锁，常被用作读共享锁；
LOCK_EX：排他锁，同时只允许一个进程使用，常被用作写锁；
LOCK_UN：释放锁。

### tidwall/redcon

### spaolacci/murmur3
+ golang哈希算法性能对比md5,crc32,sha1,murmur3: <https://blog.csdn.net/raoxiaoya/article/details/127411731>
+ murmur3要比md5和sha1要快很三倍多。


### server端
+ 处理命令: cmd/cli.go execClientCommand

