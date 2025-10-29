// 协程（goroutine）是 Go 中轻量级的执行线程。

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

	// 假设我们有一个函数调用 `f(s)`，以下是同步调用的方式。
	f("direct")

	// 若想在协程中调用函数，可写成 `go f(s)`。
	// 新创建的协程会与当前协程并发执行。
	go f("goroutine")

	// 也可以对匿名函数启动协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在两个函数调用都在各自的协程中异步运行。
	// 这里通过睡眠等待它们结束（更稳妥的做法是使用 [WaitGroup](waitgroups)）。
	time.Sleep(time.Second)
	fmt.Println("done")
}
