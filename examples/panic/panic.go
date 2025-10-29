// panic 通常表示出现了意料之外的问题。
// 它常被用于对不该在正常运行中出现的错误快速失败，或当我们无法优雅处理时直接终止。

package main

import "os"

func main() {

	// 本站其他示例会用 panic 来检查意外错误。
	// 这是唯一一个故意触发 panic 的程序。
	panic("a problem")

	// panic 的常见用途是当函数返回我们无法（或不愿）处理的错误时直接终止。
	// 下例演示在创建文件时遇到意外错误就 panic。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
