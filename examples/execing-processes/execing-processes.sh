# 运行程序后会被 `ls` 进程替换。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# 需要注意，Go 并没有传统的 Unix `fork`。
# 不过通过协程、启动进程以及 `exec` 足以覆盖大多数 `fork` 的用例。
