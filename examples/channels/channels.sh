# 运行程序时，`"ping"` 消息会通过 channel 成功地在两个 goroutine 间传递。
$ go run channels.go 
ping

# 默认情况下，发送和接收都会阻塞直到双方就绪。
# 因此无需额外同步机制，就能等到 `"ping"` 被接收。
