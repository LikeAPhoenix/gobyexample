// 通道是连接并发协程的管道。
// 一个协程可以向通道发送值，另一个协程从同一通道接收值。

package main

import "fmt"

func main() {

	// 使用 `make(chan 类型)` 创建新通道。
	// 通道的类型由其传递的值决定。
	messages := make(chan string)

	// 通过 `channel <- 值` 向通道发送数据。
	// 这里在新协程中向 `messages` 通道发送 `"ping"`。
	go func() { messages <- "ping" }()

	// 使用 `<-channel` 从通道接收值。
	// 下面接收刚才发送的 `"ping"` 并打印。
	msg := <-messages
	fmt.Println(msg)
}
