// 行过滤器是一类常见程序，从标准输入读取数据、处理后将结果输出到标准输出。
// `grep`、`sed` 就是典型的行过滤器。

// 下面的 Go 示例会把输入文本转换为大写。
// 你可以参考这个模式编写自己的行过滤器。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 用缓冲扫描器包装未缓冲的 `os.Stdin`，
	// 便可使用 `Scan` 方法逐行读取输入。
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` 返回当前的词元，这里就是下一行文本。
		ucl := strings.ToUpper(scanner.Text())

		// 输出转换成大写的行。
		fmt.Println(ucl)
	}

	// `Scan` 完成后检查是否出错。
	// 文件结束不会被视作错误。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
