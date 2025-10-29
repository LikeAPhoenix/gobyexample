// 有时需要按照不同于默认顺序的规则排序。
// 例如想按字符串长度而非字母序排序。
// 下面演示如何在 Go 中自定义排序。

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// 编写一个比较字符串长度的函数，`cmp.Compare` 在这里很有帮助。
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// 使用自定义比较器调用 `slices.SortFunc`，即可按名称长度排序 `fruits`。
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// 同样的技巧也适用于非内建类型的 slice。
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// 使用 `slices.SortFunc` 按年龄排序 `people`。
	//
	// 注意：若 `Person` struct 较大，slice 可存放 `*Person` 来避免拷贝，并相应调整排序函数。
	// 若不确定影响，可进行[基准测试](testing-and-benchmarking)。
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
