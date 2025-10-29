// [命令行参数](https://en.wikipedia.org/wiki/Command-line_interface#Arguments) 是为程序传递参数的常见方式。
// 例如执行 `go run hello.go` 时，`go` 程序接收到 `run` 与 `hello.go` 两个参数。

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` 提供原始命令行参数。
	// 切片的第一个元素是可执行文件路径，而 `os.Args[1:]` 保存传入的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 可以像普通切片那样通过索引获取特定参数。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
