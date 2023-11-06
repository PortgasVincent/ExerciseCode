服务器: aws, tx, ali都有
重要项目的QPS: 几十k
不需要融资，赤城子作为投资方
前端后端分离，根据项目配置前后端qa等

公司:S

mvcc的version怎么来的 undo log作用  bin log作用?
undo log: MVCC，即多版本控制。在MySQL数据库InnoDB存储引擎中，用undo Log来实现多版本并发控制(MVCC)。当读取的某一行被其他事务锁定时，它可以从undo log中分析出该行记录以前的数据版本是怎样的，从而让用户能够读取到当前事务操作之前的数据【快照读】。
mvcc判断是否可读时，有一个指针指向之前的版本

什么隔离级别下能使用bin log？
RR(可重复读)及以上可以用statement, 以下用row
如果是RC(提交读), 用statement出现的问题是:事务A先执行，但是commit的晚，事务B后执行，commit的早，从服务器按照commit顺序执行的话就会数据不一致


redis的多路复用 redis是单线程的，为了防止单线程阻塞



跳表和树的比较
范围查找，插入和删除，查找单个key都为O(log n), 实现简单

后序遍历

zset的底层结构?
redis结构:
string: int, row 或者 embstr
    如果只是整数，用int 否则变成SDS对象(raw和embstr都是SDS对象), 只会变成raw
    长度<32字节时，用embstr, 对象和字符串放在一次性分配的一块大内存中，便于内存分配和释放
    执行任何修改命令，也会变成raw
list: ziplist 或者 linkedlist
    长度<64字节且数量<512时用ziplist (配置可更改)
hash: ziplist 或 hashtable
    元素长度都小于64字节切元素数量小于512个 (配置可更改)
set: intset 或 hashtable(字典)
    元素都是整数值且元素数量不超过512个时用整数集
zset: ziplist 或者 skiplist(字典+跳表) 
    元素数量小于128并且所有成员长度小于64字节用ziplist， 否则用dict+skiplist
    字典负责O(1) 跳表负责排序等操作 二者节点用指针指向同一个val和score