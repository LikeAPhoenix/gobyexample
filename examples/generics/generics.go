// 自 Go 1.18 起，语言新增了泛型（也称类型参数）。

package main

import "fmt"

// 下面的泛型函数 `SlicesIndex` 接受任意 `comparable` 类型的 slice 和一个元素，
// 返回该元素首次出现的索引，若不存在则返回 -1。
// `comparable` 约束表示该类型支持 `==` 与 `!=` 比较。
// 关于签名的详细说明可参考[这篇博文](https://go.dev/blog/deconstructing-type-parameters)。
// 标准库中也提供了相同功能的 [slices.Index](https://pkg.go.dev/slices#Index)。
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// 泛型类型示例：`List` 是可以存储任意类型值的单链表。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 可以像普通类型一样为泛型类型定义方法，但要保留类型参数，类型名是 `List[T]` 而非 `List`。
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// `AllElements` 将链表元素全部收集为 slice。
// 在下一个示例中会看到更惯用的遍历自定义类型的方式。
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// 调用泛型函数时通常可以依赖“类型推断”。
	// 这里无需显式写出 `S` 和 `E`，编译器会自动推断。
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// 当然也可以显式地写出类型参数。
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
