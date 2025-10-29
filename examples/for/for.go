// `for` 是 Go 唯一的循环结构。
// 下面展示几种常见的 `for` 循环写法。

package main

import "fmt"

func main() {

	// 最基础的写法，仅包含一个条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 标准的初始化/条件/后置三段式循环。
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// 还可以通过对整数使用 `range` 来实现“执行 N 次”的迭代。
	for i := range 3 {
		fmt.Println("range", i)
	}

	// 没有条件的 `for` 会一直循环，直到使用 `break` 或在外层函数中 `return`。
	for {
		fmt.Println("loop")
		break
	}

	// 也可以使用 `continue` 进入下一次迭代。
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
