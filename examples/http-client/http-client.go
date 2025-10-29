// Go 标准库的 `net/http` 包为 HTTP 客户端与服务端提供了良好支持。
// 本示例使用它发起简单的 HTTP 请求。
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// 向服务器发起 HTTP GET 请求。
	// `http.Get` 是创建 `http.Client` 并调用其 `Get` 方法的便捷封装，
	// 底层使用具有合理默认值的 `http.DefaultClient`。
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 输出 HTTP 响应状态。
	fmt.Println("Response status:", resp.Status)

	// 打印响应体的前 5 行。
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
