// Go 支持[匿名函数](https://en.wikipedia.org/wiki/Anonymous_function)，
// 它们可以构成<a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>闭包</em></a>。
// 当需要内联定义函数而无需命名时，匿名函数非常有用。

package main

import "fmt"

// 函数 `intSeq` 返回另一个函数，我们在函数体内匿名定义它。
// 返回的函数会“捕获”变量 `i`，从而形成闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// 调用 `intSeq` 并将返回的函数赋给 `nextInt`。
	// 该函数值拥有自己的 `i`，每次调用都会更新。
	nextInt := intSeq()

	// 多次调用 `nextInt` 可以看到闭包保持的状态。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 为了验证状态确实与具体函数实例绑定，再创建一个新的闭包试试。
	newInts := intSeq()
	fmt.Println(newInts())
}
