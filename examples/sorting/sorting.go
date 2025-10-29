// `slices` 包为内建类型和自定义类型提供排序功能。
// 我们先从内建类型的排序看起。

package main

import (
	"fmt"
	"slices"
)

func main() {

	// 排序函数是泛型的，适用于任意“可排序”的内建类型。
	// 可排序类型列表参见 [cmp.Ordered](https://pkg.go.dev/cmp#Ordered)。
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// 排序 `int` 的例子。
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// `slices` 包还能检查切片是否已经有序。
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
