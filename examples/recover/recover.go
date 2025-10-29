// Go 提供了内建函数 `recover`，可以从 panic 中恢复。
// `recover` 能阻止 panic 终止程序，让其继续执行。

// 实际应用中，例如服务器不希望因为某个客户端连接触发严重错误就崩溃，
// 而是应该关闭该连接并继续服务其他客户端。
// Go 的 `net/http` 默认就会这样处理 HTTP 服务器。

package main

import "fmt"

// This function panics.
func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` 必须在延迟执行的函数中调用。
	// 当外层函数发生 panic 时，该 defer 会被触发，其中的 `recover` 会捕获 panic。
	defer func() {
		if r := recover(); r != nil {
			// `recover` 的返回值即 `panic` 抛出的错误。
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// 这行不会执行，因为 `mayPanic` 会 panic。
	// `main` 会在 panic 处中断，并跳转到 defer 的闭包。
	fmt.Println("After mayPanic()")
}
