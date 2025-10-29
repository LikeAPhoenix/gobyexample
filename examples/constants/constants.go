// Go 支持字符、字符串、布尔和数值类型的常量。

package main

import (
	"fmt"
	"math"
)

// 使用 `const` 声明常量值。
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` 语句也可以出现在函数体内部。
	const n = 500000000

	// 常量表达式执行运算时具有任意精度。
	const d = 3e20 / n
	fmt.Println(d)

	// 数值常量在被显式转换之前是无类型的。
	fmt.Println(int64(d))

	// 为数值指定类型的一种方式，是把它放在需要该类型的上下文中，
	// 比如变量赋值或函数调用。例如这里的 `math.Sin` 期望获得 `float64`。
	fmt.Println(math.Sin(n))
}
