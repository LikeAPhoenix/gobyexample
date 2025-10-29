// 使用 `net/http` 包可以轻松编写基础 HTTP 服务器。
package main

import (
	"fmt"
	"net/http"
)

// `net/http` 的核心概念是 *handler*。
// handler 是实现 `http.Handler` 接口的对象。
// 常见做法是将符合签名的函数通过 `http.HandlerFunc` 适配。
func hello(w http.ResponseWriter, req *http.Request) {

	// 作为 handler 的函数会接收 `http.ResponseWriter` 与 `http.Request`。
	// 响应写入器用于构造 HTTP 响应，这里返回简单文本。
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// 这个 handler 略复杂一些：读取所有请求头并回写到响应体。
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// 使用 `http.HandleFunc` 将 handler 注册到路由。
	// 该函数会配置 `net/http` 包的默认路由器。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 最后调用 `ListenAndServe`，传入端口与 handler。
	// `nil` 表示使用刚配置的默认路由器。
	http.ListenAndServe(":8090", nil)
}
