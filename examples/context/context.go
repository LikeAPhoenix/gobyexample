// 上一示例介绍了如何搭建简单的 [HTTP 服务器](http-server)。
// HTTP 服务非常适合展示 `context.Context` 控制取消的用法。
// `Context` 可以在跨越 API 边界与协程时携带截止时间、取消信号以及请求范围内的值。
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// `net/http` 会为每个请求创建 `context.Context`，可通过 `Context()` 方法获取。
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// 在回复客户端前先等待几秒，用以模拟服务器正在处理任务。
	// 处理过程中需要关注上下文的 `Done()` 通道，一旦收到信号就应尽快取消并返回。
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// `Err()` 会返回 `Done()` 关闭的原因。
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// 和之前一样，在 `/hello` 路由注册处理器并启动服务。
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
