# 运行写文件示例。
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# 查看写入结果。
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# 接着会把这些文件 I/O 思路应用到 `stdin` 与 `stdout` 流上。
