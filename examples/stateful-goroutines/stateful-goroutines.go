// 在上一示例中，我们使用 [mutex](mutexes) 显式加锁以同步多个 goroutine 对共享状态的访问。
// 另一种实现方式是利用 goroutine 与 channel 的内建同步机制。
// 这种基于 channel 的方式符合 Go “通过通信共享内存”的理念，
// 并让每份数据都只被一个 goroutine 拥有。

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 本示例中，状态由单个 goroutine 持有，从而确保不会因并发访问而损坏。
// 其他 goroutine 若要读写状态，需要向持有者发送消息并等待回应。
// `readOp` 与 `writeOp` struct 封装了请求以及返回结果的 channel。
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 与之前一样，统计执行的操作次数。
	var readOps uint64
	var writeOps uint64

	// `reads` 与 `writes` channel 分别用于其他 goroutine 发起读写请求。
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 下方 goroutine 负责持有 `state`，与上一示例一样是一个 map，但此处仅该 goroutine 可访问。
	// 它循环 `select` 监听 `reads` 与 `writes`，接收到请求后执行操作并将结果写回 `resp`。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个 goroutine，通过 `reads` 向状态持有者发起读取。
	// 每次读取需构建 `readOp`，发送到 `reads`，再从 `resp` 接收结果。
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 对写操作也启动 10 个 goroutine，流程类似。
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让这些 goroutine 运行 1 秒。
	time.Sleep(time.Second)

	// 最后读取并输出操作次数。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
