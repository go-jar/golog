# 思路
1. 定义 LOG 的级别，LOG 中的方法，LOG 的格式，LOG 的写目的地（控制台，文件）。
2. 根据 LOG 的目的地定义如下几种形式的 Writer，均需实现 IWriter 接口：
	- 定义将 LOG 记录到控制台的 ConsoleWriter。
	- 定义将 LOG 同步记录到文件的 FileWriter。
	- 定义将 LOG 异步记录到文件的 AsyncWriter。AsyncWriter 采用装饰者模式，在 FileWriter 的基础上增加了异步操作。
3. 定义多种 LOG 格式，供选用。SimpleFormat 是一个最基本的 LOG 格式，ConsoleFormat 采用装饰者模式，在 SimpleFormat 基础上附加了一些样式。

# 参考
https://github.com/goinbox/golog
https://www.jianshu.com/p/20d0f74c3c08
