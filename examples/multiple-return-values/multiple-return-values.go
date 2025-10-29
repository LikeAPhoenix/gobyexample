// Go 原生支持多返回值。
// 这一特性在惯用 Go 中很常见，例如同时返回结果与错误值。

package main

import "fmt"

// 函数签名中的 `(int, int)` 表示该函数返回两个 `int`。
func vals() (int, int) {
	return 3, 7
}

func main() {

	// 使用多重赋值接收函数返回的两个值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 若只想要部分返回值，可以用空白标识符 `_` 忽略其他值。
	_, c := vals()
	fmt.Println(c)
}
