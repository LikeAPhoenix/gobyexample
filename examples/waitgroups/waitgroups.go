// 若要等待多个 goroutine 完成，可以使用 WaitGroup。

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// 通过睡眠模拟耗时任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// WaitGroup 用于等待这里启动的所有 goroutine 执行完毕。
	// 注意：如果要把 WaitGroup 传入函数，应该以指针方式传递。
	var wg sync.WaitGroup

	// 使用 `WaitGroup.Go` 启动多个 goroutine。
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// 阻塞等待所有 goroutine 结束；当 goroutine 调用的函数返回时，就表示已完成。
	wg.Wait()

	// 这种方式无法直接传递工作 goroutine 产生的错误。
	// 更复杂的场景可考虑使用 [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)。
}
