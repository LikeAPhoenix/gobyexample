// 使用 `os.Exit` 立即以指定状态码退出。

package main

import (
	"fmt"
	"os"
)

func main() {

	// 调用 `os.Exit` 时不会执行 `defer`，所以这行 `fmt.Println` 不会被调用。
	defer fmt.Println("!")

	// 以状态码 3 退出。
	os.Exit(3)
}

// 注意 Go 与 C 等语言不同，`main` 的返回值不是退出码。
// 如果需要非零退出码，应调用 `os.Exit`。
