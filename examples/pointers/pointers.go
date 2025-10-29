// Go 支持<a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)"><em>指针</em></a>，
// 这使得我们可以在程序中传递值或结构的引用。

package main

import "fmt"

// 我们用两个函数 `zeroval` 和 `zeroptr` 对比值传递与指针的差异。
// `zeroval` 的参数是 `int`，因此实参按值传递，它得到的是调用者 `ival` 的副本。
func zeroval(ival int) {
	ival = 0
}

// 相对地，`zeroptr` 的参数类型为 `*int`，表示接收 `int` 指针。
// 函数体内的 `*iptr` 会解引用该指针，从地址取得当前的值。
// 给解引用的指针赋值就会修改该地址上的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// `&i` 语法会取得 `i` 的内存地址，也就是指向 `i` 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针本身也可以被打印出来。
	fmt.Println("pointer:", &i)
}
