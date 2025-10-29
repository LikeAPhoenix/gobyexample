# 程序运行后会阻塞等待信号。
# 在终端按 `Ctrl+C`（显示为 `^C`）即可发送 `SIGINT`，程序会打印 `interrupt` 并退出。
$ go run signals.go
awaiting signal
^C
interrupt
exiting
