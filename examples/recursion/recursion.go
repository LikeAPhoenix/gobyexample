// Go 支持<a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>递归函数</em></a>。
// 这是一个经典示例。

package main

import "fmt"

// `fact` 函数会递归调用自身，直到遇到基线条件 `fact(0)`。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// 匿名函数同样可以递归，不过需要先用 `var` 声明变量存放函数，再进行定义。
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// 由于 `fib` 在 `main` 中已声明，Go 可以识别这里调用的是同一个函数。
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
