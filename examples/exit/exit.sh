# 使用 `go run` 执行时，`go` 工具会捕获并打印退出状态。
$ go run exit.go
exit status 3

# 若编译后直接运行可执行文件，可以在终端看到退出码。
$ go build exit.go
$ ./exit
$ echo $?
3

# 注意程序中的 `!` 从未打印出来。
