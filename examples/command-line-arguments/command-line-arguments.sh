# 测试命令行参数时，先使用 `go build` 构建可执行文件。
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# 接下来介绍更高级的命令行处理：标志参数。
