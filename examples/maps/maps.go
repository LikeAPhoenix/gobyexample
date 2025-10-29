// map 是 Go 内建的[关联数据类型](https://en.wikipedia.org/wiki/Associative_array)，
// 在其他语言中有时被称作 hash 或 dict。

package main

import (
	"fmt"
	"maps"
)

func main() {

	// 使用内建的 `make` 可以创建空 map：`make(map[key-type]val-type)`。
	m := make(map[string]int)

	// 使用常见的 `name[key] = val` 语法设置键值对。
	m["k1"] = 7
	m["k2"] = 13

	// 使用 `fmt.Println` 打印 map 会展示其中所有键值对。
	fmt.Println("map:", m)

	// 通过 `name[key]` 获取对应的值。
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 如果键不存在，会返回该值类型的[零值](https://go.dev/ref/spec#The_zero_value)。
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// 对 map 调用内建的 `len` 可以得到键值对数量。
	fmt.Println("len:", len(m))

	// 内建函数 `delete` 可用来删除键值对。
	delete(m, "k2")
	fmt.Println("map:", m)

	// 若要删除所有键值对，可以使用内建的 `clear`。
	clear(m)
	fmt.Println("map:", m)

	// 从 map 读取时的可选第二个返回值表示键是否存在。
	// 这样就能区分键缺失和键存在但值为零值（例如 `0` 或 `""`）的情况。
	// 在这里我们不需要实际的值，因此使用空白标识符 `_` 丢弃它。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 还可以使用一行语法声明并初始化新的 map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// `maps` 包中提供了很多适用于 map 的实用函数。
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
