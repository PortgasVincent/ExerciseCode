逻辑题: D B B E A
Golang:
1.
1 2

2.
AB

3.
B

4.
A

5.
B

网络，操作系统，数据库
1.
3 4

2.
Session的生命周期只有本次会话，Cookie生命周期取决于过期时间

3.
400: 无效请求 500: 服务器问题

4.
grep: 在文件中查找字符串   top: 显示进程状况
find: 查找文件      MV: 移动文件到目标目录
cat: 输出文件内容    du: 文件占用空间
wc:               chmod: 改权限相关
ps: 显示进程详细内容

5.
代表三种权限  r: 读 w: 写 x: 执行   目录

6.
模拟sql执行，展示执行细节（表，顺序，调用行等），常用来优化sql

列出运行的进程

7.
单例模式，工厂模式

算法：
1.
// 绝对值最小的数可能有两个，输出其中一个
func Find(arr []int) int {
	l := len(arr)
	if l == 0 {
		return 0
	}
	if l == 1 || arr[0] >= 0 {
		return arr[0]
	}
	return find(arr, 0, l-1)
}

func find(arr []int, left, right int) int {
	if left == right {
		return arr[left]
	}
	if right-left == 1 {
		return min(arr[left], arr[right])
	}
	mid := (left + right) / 2
	if arr[mid] >= 0 {
		return find(arr, left, mid)
	}
	return find(arr, mid, right)
}

func min(a, b int) int {
	if math.Abs(float64(a)) < math.Abs(float64(b)) {
		return a
	}
	return b
}

2.
func step(k int) int {
	if k < 3 {
		return 1
	}
	if k == 3 {
		return 2
	}
	return step(k-1) + step(k-3)
}
