，  go module
http://go-day.cn/detail/6.html
版本控制：go path -> vendor -> go module

mysql version特性 8.0
https://www.cnblogs.com/bantiaoxianyu/p/16485658.html
更好的对Nosql和Json支持，移除查询缓存

go version特性
1.11 支持go modules模式 增加新的依赖管理模块并且更加易于管理项目中所需要的模块。模块是存储在文件树中的 Go 包的集合，其根目录中包含 go.mod 文件
从 Go 1.13/1.16？ 开始，模块模式将成为默认模式。
1.18 支持泛型
1.20 支持slice向数组转化
1.21 降低GC尾部延迟


常用的包 rand
func Intn(n int) int // return [0,n)

go控制并发，控制异步
sync.waitgroup, channel, context，time.timer
异步不需要读取结果可以用不同的goroutine，用channel通信等

context直接取消和等一段时间取消?
http://www.flysnow.org/2017/05/12/go-in-action-go-context.html
type Context interface {
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key interface{}) interface{}
}
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context

常用的mysql类型，varchar和text区别   text能为空吗?
https://stackoverflow.com/questions/25300821/difference-between-varchar-and-text-in-mysql
TEXT:
fixed max size of 65535 characters (you cannot limit the max size)
takes 2 + c bytes of disk space, where c is the length of the stored string.
cannot be (fully) part of an index. One would need to specify a prefix length.
TEXT does not support default values of anything but NULL!!! text默认值只能为nul

VARCHAR(M):
variable max size of M characters
M needs to be between 1 and 65535
takes 1 + c bytes (for M ≤ 255) or 2 + c (for 256 ≤ M ≤ 65535) bytes of disk space where c is the length of the stored string
can be part of an index


text和blog区别?
https://www.cnblogs.com/printN/p/7463737.html
TEXT与BLOB的主要差别就是BLOB保存二进制数据，TEXT保存字符数据

从1-100取三个不重复的随机数？
黑名单映射 https://leetcode.cn/problems/random-pick-with-blacklist/description/
