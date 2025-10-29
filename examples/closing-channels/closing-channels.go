// 关闭 channel 意味着不会再有新的值发送到该 channel。
// 这通常用于向接收方告知处理完成。

package main

import "fmt"

// 本示例使用 `jobs` channel 把工作从 `main()` goroutine 传递给工作 goroutine。
// 当没有更多任务时，就关闭 `jobs` channel。
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 工作 goroutine 通过 `j, more := <-jobs` 持续接收任务。
	// 在这种双返回值形式中，若 `jobs` 已关闭且 channel 中没有值，`more` 会是 `false`。
	// 利用这一点，在任务全部完成时通过 `done` 通知主 goroutine。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 向工作 goroutine 发送 3 个任务，然后关闭 channel。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 使用之前的[同步](channel-synchronization)技巧等待工作 goroutine 结束。
	<-done

	// 从关闭的 channel 读取会立即成功，并返回元素类型的零值。
	// 第二个可选返回值为 `true` 表示读取到的是正常发送的数据；
	// 为 `false` 表示由于 channel 已关闭且为空，因此返回零值。
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
