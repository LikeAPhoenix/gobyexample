// 在之前的示例中，我们用[原子操作](atomic-counters)管理简单计数器。
// 如果状态更复杂，可以使用[互斥锁](https://en.wikipedia.org/wiki/Mutual_exclusion)在多协程间安全访问数据。

package main

import (
	"fmt"
	"sync"
)

// `Container` 保存了一组计数器。
// 因为需要被多个协程并发更新，所以加上 `Mutex` 同步访问。
// 注意互斥锁不能被复制，因此该结构体若需要传递，应以指针方式进行。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// 访问 `counters` 前锁住互斥锁，并在函数结束时通过 [defer](defer) 解锁。
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// 互斥锁的零值即可直接使用，因此无需显式初始化。
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// 该函数循环递增指定名字的计数器。
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// 并发启动多个协程，它们共享同一个 `Container`，
	// 其中两个协程还会访问同一个计数器。
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// 等待所有协程完成。
	wg.Wait()
	fmt.Println(c.counters)
}
