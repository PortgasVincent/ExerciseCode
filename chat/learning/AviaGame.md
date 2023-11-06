AviaGame: 主要是海外真金竞技游戏，非博彩

GMP中，四核cpu对应几个线程？ 四个
观察程序CPU切换？哪些操作比较消耗CPU？
pprof观察
从生成的svg图中可以看到，mallocgc占用时长较多。原因是使用堆分配内存太频繁细碎导致的。其中，newobject（对应new关键字）、makeslice（对应make关键字）、growslice（对应append关键字）调用过于频繁，最终导致mallocgc占用时间非常长。
https://zhuanlan.zhihu.com/p/157326385