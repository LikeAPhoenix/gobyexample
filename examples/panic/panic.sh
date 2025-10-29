# 运行该程序会触发 panic，打印错误信息与 goroutine 调用栈，并以非零状态退出。

# `main` 中首次 panic 发生时，程序立即退出，后续代码不会执行。
# 若想观察创建临时文件的部分，只需注释掉第一个 panic。
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# 请注意，与某些大量依赖异常处理错误的语言不同，Go 更习惯使用返回值来指示错误。
