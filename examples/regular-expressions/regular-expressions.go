// Go 内建支持[正则表达式](https://en.wikipedia.org/wiki/Regular_expression)。
// 下面演示正则相关的一些常见操作。

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// 判断某个模式是否匹配给定字符串。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面直接使用了字符串模式，但在其他场景通常需要先 `Compile`
	// 得到优化后的 `Regexp` struct。
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 该 struct 上提供了许多方法。
	// 下面再次执行与之前相同的匹配测试。
	fmt.Println(r.MatchString("peach"))

	// 查找第一个匹配。
	fmt.Println(r.FindString("peach punch"))

	// 同样查找第一次匹配，但返回匹配的起止索引而非文本。
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// `Submatch` 变体会同时返回整个模式匹配及子匹配的信息。
	// 比如下例会返回 `p([a-z]+)ch` 及 `([a-z]+)` 的结果。
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 这一版本则返回匹配及子匹配的索引信息。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// `All` 变体会匹配输入中的所有结果，而非仅第一个。
	// 下面示例查找所有匹配项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 其他方法也有对应的 `All` 版本。
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// 将第二个参数设为非负整数即可限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 以上示例使用字符串参数，因此调用的是 `MatchString` 这类方法。
	// 也可以传递 `[]byte`，此时方法名去掉 `String` 后缀即可。
	fmt.Println(r.Match([]byte("peach")))

	// 创建正则表达式的全局变量时可以使用 `MustCompile`。
	// 它会在发生错误时 panic，而不是返回错误值，因此适合用于确保初始化成功。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// `regexp` 包也可以用来替换字符串中的匹配部分。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `ReplaceAllFunc` 变体允许通过自定义函数转换匹配文本。
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
