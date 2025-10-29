// 可以利用通道在协程之间同步执行。
// 下面示例通过阻塞接收等待协程结束。
// 如果要等待多个协程完成，可以考虑使用 [WaitGroup](waitgroups)。

package main

import (
	"fmt"
	"time"
)

// 这是将在协程中运行的函数。
// `done` 通道用于通知其他协程该函数已完成工作。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值表示已完成。
	done <- true
}

func main() {

	// 启动工作协程，并传入用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)

	// 阻塞等待，直到收到工作协程的通知。
	<-done
}
