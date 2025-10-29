// `range` 可以遍历多种内建数据结构。
// 下面回顾几种常见用法。

package main

import "fmt"

func main() {

	// 使用 `range` 累加切片中的数字，数组也可以这样操作。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 遍历数组或切片时，`range` 会同时返回索引和值。
	// 若不需要索引可以用空白标识符 `_` 忽略；有时也会用到索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 遍历 map 时，`range` 返回键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 也可以只遍历 map 的键。
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 对字符串使用 `range` 会按 Unicode 码点遍历。
	// 第一个返回值是该 `rune` 的起始字节位置，第二个是 `rune` 本身。
	// 更多细节见[字符串与 rune](strings-and-runes)。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
