// [环境变量](https://en.wikipedia.org/wiki/Environment_variable) 是向 Unix 程序传递配置的通用机制（参见[十二要素应用](https://www.12factor.net/config)）。
// 下面演示如何设置、读取与列出环境变量。

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 使用 `os.Setenv` 设置键值对，`os.Getenv` 获取指定键的值。
	// 如果环境中不存在该键，会返回空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// `os.Environ` 会返回所有环境变量的列表，以 `KEY=value` 形式组成字符串切片。
	// 可以用 `strings.SplitN` 拆分键和值。
	// 此处仅打印所有键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
