// 函数在 Go 中占据核心地位。
// 下面通过几个例子来了解函数的用法。

package main

import "fmt"

// 这是一个接受两个 `int` 并返回它们之和的函数。
func plus(a int, b int) int {

	// Go 需要显式返回值，不会自动返回最后一个表达式。
	return a + b
}

// 当多个参数类型相同时，可以省略前面参数的类型，只在最后一个参数上写一次。
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// 使用熟悉的 `name(args)` 方式调用函数即可。
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
