// [命令行标志](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option) 是命令行程序指定选项的常用方式。
// 例如在 `wc -l` 中，`-l` 就是一个命令行标志。

package main

// Go 提供了 `flag` 包来解析命令行标志。
// 我们将用它构建示例程序。
import (
	"flag"
	"fmt"
)

func main() {

	// 可以声明字符串、整数、布尔等基础类型标志。
	// 这里声明字符串标志 `word`，默认值为 `"foo"` 并附带说明。
	// `flag.String` 返回的是字符串指针，稍后会用到。
	wordPtr := flag.String("word", "foo", "a string")

	// 以下以相同方式声明 `numb` 与 `fork` 标志。
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// 还可以把标志值绑定到已有变量。
	// 此时需要把变量指针传给声明函数。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 所有标志声明完成后，调用 `flag.Parse()` 执行解析。
	flag.Parse()

	// 最后输出解析结果及剩余的位置参数。
	// 注意需要解引用如 `*wordPtr` 才能获得实际值。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
