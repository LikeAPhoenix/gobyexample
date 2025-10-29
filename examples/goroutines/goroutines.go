// goroutine 是 Go 中轻量级的执行线程。

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设我们有一次函数调用 `f(s)`，下面是常规的同步调用方式。
	f("direct")

	// 若要在 goroutine 中调用该函数，使用 `go f(s)`。
	// 新创建的 goroutine 会与当前 goroutine 并发执行。
	go f("goroutine")

	// 也可以为一次匿名函数调用启动 goroutine。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个函数调用分别在独立的 goroutine 中异步运行。
	// 这里通过短暂休眠等待它们结束（更稳妥的方式是使用 [WaitGroup](waitgroups)）。
	time.Sleep(time.Second)
	fmt.Println("done")
}
