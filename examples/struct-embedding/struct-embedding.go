// Go 支持 struct 与接口的嵌入，用以表达更加自然的类型组合。
// 不要将其与 Go 1.16+ 引入的 [`//go:embed`](embed-directive) 混淆，后者用于将文件目录嵌入二进制。

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` 内嵌了 `base`，看起来像一个没有字段名的字段。
type container struct {
	base
	str string
}

func main() {

	// 使用字面量创建 struct 时，需要显式初始化被嵌入的字段，
	// 此时嵌入类型本身充当字段名。
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// 可以直接在 `co` 上访问嵌入的字段，例如 `co.num`。
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// 也可以写出完整路径，借助嵌入类型名访问。
	fmt.Println("also num:", co.base.num)

	// `container` 嵌入了 `base`，因此 `base` 的方法同样可在 `container` 上调用。
	// 这里直接在 `co` 上调用来自 `base` 的方法。
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// 嵌入带方法的 struct 可以把这些接口实现“继承”给外层结构。
	// 因为 `container` 嵌入了 `base`，它也实现了 `describer` 接口。
	var d describer = co
	fmt.Println("describer:", d.describe())
}
