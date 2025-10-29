# 运行程序时，会先看到同步调用的输出，然后是两个协程的结果。
# 因为协程由 Go 运行时并发调度，它们的输出可能交错。
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# 接下来看看并发编程中与协程相辅相成的概念：通道。
