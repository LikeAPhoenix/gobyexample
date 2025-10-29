// Go 继承了传统 `printf` 风格的强大字符串格式化能力。
// 下面展示一些常见用法。

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go 提供了多种格式化动词（verb）用于输出不同的值。
	// 例如这里输出 `point` struct 的一个实例。
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// 对于 struct，`%+v` 会额外包含字段名字。
	fmt.Printf("struct2: %+v\n", p)

	// `%#v` 会打印该值在 Go 语法中的表示，也就是能生成此值的源代码片段。
	fmt.Printf("struct3: %#v\n", p)

	// `%T` 用来打印值的类型。
	fmt.Printf("type: %T\n", p)

	// 布尔值的格式化非常直接。
	fmt.Printf("bool: %t\n", true)

	// 格式化整数的选项很多。
	// `%d` 表示十进制输出。
	fmt.Printf("int: %d\n", 123)

	// 输出二进制表示。
	fmt.Printf("bin: %b\n", 14)

	// 输出与给定整数对应的字符。
	fmt.Printf("char: %c\n", 33)

	// `%x` 以十六进制编码输出。
	fmt.Printf("hex: %x\n", 456)

	// 浮点数也有多种格式化方式。
	// `%f` 使用基本的十进制格式。
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` 和 `%E` 分别以科学计数法输出（大小写略有差异）。
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// `%s` 用于基本的字符串输出。
	fmt.Printf("str1: %s\n", "\"string\"")

	// `%q` 会像 Go 源码那样给字符串加双引号。
	fmt.Printf("str2: %q\n", "\"string\"")

	// 类似前面的整数示例，`%x` 会把字符串的每个字节以两个十六进制字符表示。
	fmt.Printf("str3: %x\n", "hex this")

	// `%p` 用于打印指针的表示。
	fmt.Printf("pointer: %p\n", &p)

	// 格式化数字时通常需要控制宽度与精度。
	// 要指定整数宽度，可以在动词后加数字；默认右对齐并用空格填充。
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// 浮点数同样可以指定宽度。
	// 通常会结合 `宽度.精度` 的语法，同时限制小数位数。
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// 使用 `-` 标志可以左对齐。
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 格式化字符串时也可以控制宽度，便于在表格类的输出中对齐。
	// 默认同样是右对齐。
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// 与数字一样，使用 `-` 标志实现左对齐。
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// 目前为止我们使用的是 `Printf`，它会把结果输出到 `os.Stdout`。
	// `Sprintf` 则仅格式化后返回字符串而不打印。
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// 使用 `Fprintf` 可以把格式化结果写到除了 `os.Stdout` 之外的任意 `io.Writer`。
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
