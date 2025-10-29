// 本示例演示如何使用 goroutine 和 channel 实现工作池。

package main

import (
	"fmt"
	"time"
)

// 定义工作 goroutine，我们会同时运行多个实例。
// 工作 goroutine 从 `jobs` channel 接收任务，并把结果发送到 `results`。
// 通过每个任务睡眠 1 秒来模拟耗时操作。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// 要使用工作池，需要将任务发送给它们并收集结果，因此创建两个 channel。
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动 3 个工人 goroutine，此时尚未有任务，因此它们会阻塞等待。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送 5 个任务后关闭 channel，表示任务已全部分发。
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后收集所有结果，这也能确保所有工作 goroutine 都执行完毕。
	// 另一种等待多个 goroutine 完成的方法是使用 [WaitGroup](waitgroups)。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
