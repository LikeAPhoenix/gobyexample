// 在[前面的示例](range-over-built-in-types)中我们看到 `for`/`range` 遍历基础数据结构。
// 同样可以用它来遍历 channel 中接收到的值。

package main

import "fmt"

func main() {

	// 这里准备遍历 `queue` channel 中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// `range` 会随着 `queue` 中的值被接收而遍历。
	// channel 在上方已关闭，因此接收完两个元素后迭代就结束。
	for elem := range queue {
		fmt.Println(elem)
	}
}
