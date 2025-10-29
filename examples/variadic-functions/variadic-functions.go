// [可变参数函数](https://en.wikipedia.org/wiki/Variadic_function) 可以接收任意数量的末尾参数。
// 例如 `fmt.Println` 就是一种常见的可变参数函数。

package main

import "fmt"

// 下面的函数可以接收任意数量的 `int` 参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// 在函数体内部，`nums` 的类型等同于 `[]int`。
	// 因此可以对它调用 `len(nums)`、使用 `range` 遍历等。
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 调用可变参数函数时，只需像普通函数那样传入各个参数。
	sum(1, 2)
	sum(1, 2, 3)

	// 如果参数已经存在于 slice 中，可以通过 `func(slice...)` 的语法展开传入。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
