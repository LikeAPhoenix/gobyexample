// 切片是 Go 中非常重要的数据类型，相比数组提供了更灵活的序列接口。

package main

import (
	"fmt"
	"slices"
)

func main() {

	// 与数组不同，切片的类型只由元素类型决定，不包含长度信息。
	// 未初始化的切片等于 nil，长度为 0。
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// 使用内建的 `make` 可以创建非空的切片。
	// 这里创建一个长度为 3 的字符串切片，初始元素都是零值。
	// 新切片默认容量等于长度；若提前知道需要扩容，可以额外传入容量。
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// 切片的取值与赋值方式与数组一致。
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` 会返回切片的长度。
	fmt.Println("len:", len(s))

	// 除了这些基础操作，切片还支持更多功能，使其比数组更强大。
	// 例如内建的 `append` 会返回追加了新元素的切片。
	// 需要注意 `append` 可能返回新的底层数组，因此要接收其返回值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片也可以使用 `copy` 复制。
	// 这里创建了一个与 `s` 长度相同的空切片 `c`，然后把 `s` 的内容复制进去。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片支持 `slice[low:high]` 这样的“切片”操作。
	// 例如下面截取的就是元素 `s[2]`、`s[3]` 和 `s[4]`。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这一步会截取到 `s[5]`（不含）。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 而这里则从 `s[2]`（含）开始截取到末尾。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 也可以在一行内声明并初始化切片变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 标准库中的 `slices` 包提供了许多切片相关的实用函数。
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// 切片可以组合成多维数据结构，内部切片的长度可以各不相同，这一点不同于多维数组。
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
