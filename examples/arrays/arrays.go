// 在 Go 中，数组是一段具有固定长度的编号元素序列。
// 在日常 Go 代码中，[切片](slices)更常见，数组则适用于某些特殊场景。

package main

import "fmt"

func main() {

	// 这里创建了一个能存放 5 个 `int` 的数组 `a`。
	// 元素类型与长度共同决定数组的类型。
	// 默认情况下数组使用零值初始化，对 `int` 来说就是 `0`。
	var a [5]int
	fmt.Println("emp:", a)

	// 使用 `array[index] = value` 可以为某个索引赋值，`array[index]` 则用于取值。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 内建函数 `len` 会返回数组长度。
	fmt.Println("len:", len(a))

	// 使用如下语法可以在一行内声明并初始化数组。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 也可以使用 `...` 让编译器推导元素个数。
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 当通过 `:` 指定索引时，中间未显式赋值的元素会被填充为零值。
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// 数组类型本身是一维的，但可以组合来构建多维数据结构。
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 也可以在声明时直接初始化多维数组。
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
