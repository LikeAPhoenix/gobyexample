// 我们常常需要在未来某个时间点执行代码，或按固定间隔重复执行。
// Go 内建的计时器（timer）与滴答器（ticker）能轻松实现这些需求。
// 本节先介绍计时器，下一节再看 [ticker](tickers)。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 计时器表示未来的一次事件。
	// 指定等待时长后，它会在对应时刻通过通道发送通知。
	// 此处的计时器等待 2 秒。
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` 会阻塞，直到通道 `C` 收到计时完成的信号。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 如果只是等待，也可以直接用 `time.Sleep`。
	// 但计时器的优势之一是可以在触发前取消，这里演示如何取消。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 额外等待一会儿，确认 `timer2` 确实已停止。
	time.Sleep(2 * time.Second)
}
