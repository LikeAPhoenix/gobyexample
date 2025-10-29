// channel 是连接并发 goroutine 的管道。
// 一个 goroutine 可以向 channel 发送值，另一个 goroutine 从同一 channel 接收值。

package main

import "fmt"

func main() {

	// 使用 `make(chan 类型)` 创建新 channel。
	// channel 的类型由其传递的值决定。
	messages := make(chan string)

	// 通过 `channel <- 值` 向 channel 发送数据。
	// 这里在新 goroutine 中向 `messages` channel 发送 `"ping"`。
	go func() { messages <- "ping" }()

	// 使用 `<-channel` 从 channel 接收值。
	// 下面接收刚才发送的 `"ping"` 并打印。
	msg := <-messages
	fmt.Println(msg)
}
