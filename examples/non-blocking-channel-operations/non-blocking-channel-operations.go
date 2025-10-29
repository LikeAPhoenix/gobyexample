// channel 的基本发送与接收都是阻塞的。
// 不过可以借助带 `default` 分支的 `select` 实现非阻塞的发送、接收，甚至多路选择。

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 非阻塞接收示例：若 `messages` 中有值，`select` 会选择 `<-messages` 分支；
	// 否则立即走 `default`。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞发送同理。
	// 当前 `messages` 无缓冲且没有接收者，因此发送失败并执行 `default`。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 也可以在 `default` 前放多个分支，实现多路非阻塞 `select`。
	// 这里尝试在 `messages` 与 `signals` 上非阻塞接收。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
