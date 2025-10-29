// 有时需要 Go 程序优雅地处理[Unix 信号](https://en.wikipedia.org/wiki/Unix_signal)。
// 例如服务器收到 `SIGTERM` 时平滑关闭，命令行工具接收 `SIGINT` 时停止处理输入。
// 下面演示如何借助 channel 处理信号。

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go 的信号通知通过在 channel 上发送 `os.Signal` 值实现。
	// 创建一个带缓冲的 channel 用于接收通知。
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` 会让指定 channel 接收所列信号。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 虽然可以在 `main` 中直接接收信号，
	// 这里改在单独 goroutine 中处理，更贴近日常的优雅关闭场景。
	done := make(chan bool, 1)

	go func() {
		// 该 goroutine 阻塞等待信号，收到后打印并通知主程序可以退出。
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 主程序在此等待，直到收到 goroutine 发来的完成信号，再退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
