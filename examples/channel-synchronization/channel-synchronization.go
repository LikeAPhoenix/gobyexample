// 可以利用 channel 在 goroutine 之间同步执行。
// 下面示例通过阻塞接收等待 goroutine 结束。
// 如果要等待多个 goroutine 完成，可以考虑使用 [WaitGroup](waitgroups)。

package main

import (
	"fmt"
	"time"
)

// 这是将在 goroutine 中运行的函数。
// `done` channel 用于通知其他 goroutine 该函数已完成工作。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值表示已完成。
	done <- true
}

func main() {

	// 启动工作 goroutine，并传入用于通知的 channel。
	done := make(chan bool, 1)
	go worker(done)

	// 阻塞等待，直到收到工作 goroutine 的通知。
	<-done
}
