// Go 管理状态的主要方式是通过通道通信，例如[工作池](worker-pools)。
// 当然也有其他手段，这里介绍如何用 `sync/atomic` 实现多协程访问的原子计数器。

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// 通过原子整型表示计数器（始终为非负）。
	var ops atomic.Uint64

	// 使用 WaitGroup 等待所有协程结束。
	var wg sync.WaitGroup

	// 启动 50 个协程，每个协程递增计数器 1000 次。
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// 使用 `Add` 以原子方式递增。
				ops.Add(1)
			}
		})
	}

	// 等待所有协程完成。
	wg.Wait()

	// 此时没有协程再写入 `ops`，但即便并发更新仍在进行，`Load` 也能安全读取当前值。
	fmt.Println("ops:", ops.Load())
}
