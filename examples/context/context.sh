# 在后台启动服务器。
$ go run context.go &

# 模拟客户端请求 `/hello`，启动后很快按 Ctrl+C 触发取消信号。
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
