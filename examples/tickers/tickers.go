// [计时器](timers) 适用于未来某个时间执行一次，
// 而 ticker 用于按固定间隔重复执行。
// 下面演示一个周期触发的 ticker，并在稍后停止它。

package main

import (
	"fmt"
	"time"
)

func main() {

	// ticker 与 timer 的机制类似，都会通过 channel 发送信号。
	// 这里每 500ms 收到一个值，并使用 `select` 等待它们。
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// ticker 也可以像 timer 一样停止。
	// 一旦停止，其 channel 就不会再收到值。
	// 此处在 1600ms 后停止。
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
