// Go 提供多种值类型，包括字符串、整数、浮点数、布尔值等。
// 下面展示几个基础示例。

package main

import "fmt"

func main() {

	// 字符串可以用 `+` 进行拼接。
	fmt.Println("go" + "lang")

	// 整数与浮点数。
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 布尔值以及常见的布尔运算。
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
