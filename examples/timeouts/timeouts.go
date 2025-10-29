// 当程序需要访问外部资源或限制执行时间时，超时控制非常重要。
// 借助通道和 `select`，在 Go 中实现超时既简单又优雅。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 假设执行某个外部调用，它会在 2 秒后把结果写入 `c1`。
	// 这里的通道带缓冲，因此协程中的发送不会阻塞。
	// 这是防止通道无人读取导致协程泄漏的常见做法。
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// 下方的 `select` 实现了超时逻辑。
	// `res := <-c1` 等待结果，而 `<-time.After` 会在 1 秒后收到超时信号。
	// `select` 会选择最先就绪的分支，因此如果操作超过 1 秒，就会执行超时分支。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 如果把超时时间延长至 3 秒，就能成功接收到 `c2` 的结果。
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
