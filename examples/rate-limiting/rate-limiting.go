// [限流](https://en.wikipedia.org/wiki/Rate_limiting) 是控制资源使用、
// 保持服务质量的重要机制。
// Go 可以借助 goroutine、channel 与 [ticker](tickers) 优雅地实现限流。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 先看最基础的限流。
	// 假设我们要限制处理传入请求的速率，可将请求放入 channel。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// `limiter` channel 每隔 200 毫秒收到一个值，充当限流器。
	limiter := time.Tick(200 * time.Millisecond)

	// 在处理每个请求前先从 `limiter` 读取，就能把速率限制为 200 毫秒一个请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 如果想在总体限流的同时允许短暂“突发（burst）”，可以给限流器加缓冲。
	// `burstyLimiter` 允许最多 3 个事件的 burst。
	burstyLimiter := make(chan time.Time, 3)

	// 预先填满 channel，表示允许的初始突发额度。
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// 之后每隔 200 毫秒向 `burstyLimiter` 追加一个值，最多保持 3 个。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 模拟再来 5 个请求，其中前 3 个会利用突发配额立即得到处理。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
