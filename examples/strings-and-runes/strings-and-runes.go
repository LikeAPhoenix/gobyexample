// Go 中的字符串是只读的字节切片。
// 语言本身和标准库都会将字符串视作以 [UTF-8](https://en.wikipedia.org/wiki/UTF-8) 编码的文本容器。
// 在其他语言里，字符串通常由“字符”组成。
// 在 Go 中，这一概念称为 `rune`，它是表示 Unicode 码点的整数。
// [这篇 Go 博客文章](https://go.dev/blog/strings) 是一个很好的入门材料。

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` 是一个字符串字面量，对应泰语中的“你好”。
	// Go 的字符串字面量使用 UTF-8 编码。
	const s = "สวัสดี"

	// 由于字符串等价于 `[]byte`，此调用会得到底层原始字节的长度。
	fmt.Println("Len:", len(s))

	// 对字符串进行索引会返回该位置的原始字节值。
	// 下面的循环输出组成 `s` 的所有字节的十六进制表示。
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// 若要统计字符串中包含多少个 rune，可以使用 `utf8` 包。
	// 注意 `RuneCountInString` 的运行时间取决于字符串长度，
	// 因为它需要依次解码每一个 UTF-8 rune。
	// 一些泰文字符由多个字节的 UTF-8 码点组成，所以统计结果可能出乎意料。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// `range` 循环会特殊处理字符串，为每个 `rune` 解码并返回它在字符串中的偏移。
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// 也可以显式调用 `utf8.DecodeRuneInString` 来完成同样的遍历。
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// 这里演示了如何把 `rune` 值传给函数。
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// 用单引号括起来的值是 rune 字面量，可以直接与 `rune` 值比较。
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
