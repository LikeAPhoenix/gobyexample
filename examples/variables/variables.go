// 在 Go 中，变量需要显式声明，编译器会借此检查函数调用等场景的类型正确性。

package main

import "fmt"

func main() {

	// 使用 `var` 可以声明一个或多个变量。
	var a = "initial"
	fmt.Println(a)

	// 也可以一次声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 会依据初始化表达式推断变量类型。
	var d = true
	fmt.Println(d)

	// 未显式初始化的变量会获得“零值”（zero value）。例如 `int` 的零值是 `0`。
	var e int
	fmt.Println(e)

	// `:=` 语法是用于声明并初始化变量的简写，此处等价于 `var f string = "apple"`。
	// 这种语法仅能在函数内部使用。
	f := "apple"
	fmt.Println(f)
}
