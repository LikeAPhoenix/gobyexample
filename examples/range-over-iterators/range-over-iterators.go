// 自 Go 1.23 起，语言支持迭代器（[iterators](https://go.dev/blog/range-functions)），
// 可以对几乎任何东西执行 range 遍历。

package main

import (
	"fmt"
	"iter"
	"slices"
)

// 再看[前一个示例](generics)中的 `List` 类型。
// 当时通过 `AllElements` 返回所有元素组成的 slice。
// 借助 Go 的迭代器，可以用更优雅的方式处理。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// `All` 返回一个迭代器，在 Go 中它是具有[特定签名](https://pkg.go.dev/iter#Seq)的函数。
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// 迭代器函数会接收另一个函数（惯例命名为 `yield`，实际可自定），
		// 并为每个元素调用 `yield`；一旦返回值为 false 就提前终止。
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// 迭代并不需要底层数据结构，甚至可以是无限序列。
// 下例生成斐波那契数列，只要 `yield` 返回 true 就会继续。
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// 因为 `List.All` 返回迭代器，可以直接在 `range` 中使用。
	for e := range lst.All() {
		fmt.Println(e)
	}

	// [slices](https://pkg.go.dev/slices) 等包提供了许多处理迭代器的工具函数。
	// 例如 `Collect` 能将任意迭代器的值收集为 slice。
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// 当循环遇到 `break` 或提前返回时，传入迭代器的 `yield` 会返回 false。
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
