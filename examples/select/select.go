// Go 的 `select` 可用于同时等待多个通道操作。
// 将协程、通道与 `select` 结合是 Go 的强大特性之一。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 示例中会在两个通道上进行选择。
	c1 := make(chan string)
	c2 := make(chan string)

	// 每个通道都会在一段时间后收到值，用于模拟阻塞的 RPC 等操作在协程中执行。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 使用 `select` 同时等待两个值，并在它们到达时打印。
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
