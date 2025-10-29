// Go 标准库提供了便捷的日志工具。
// [log](https://pkg.go.dev/log) 包适合自由格式输出，
// [log/slog](https://pkg.go.dev/log/slog) 则用于结构化输出。
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// 直接调用 `log.Println` 等函数时会使用标准 logger（日志器），
	// 它默认将合理的日志输出到 `os.Stderr`。
	// `Fatal*`、`Panic*` 等方法会在记录后终止程序。
	log.Println("standard logger")

	// 可以通过 flags 修改日志输出格式。
	// 标准 logger 默认启用 `log.Ldate` 与 `log.Ltime`，即 `log.LstdFlags`。
	// 例如我们可以额外启用微秒级时间戳。
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// 还可以输出调用 `log` 的文件名与行号。
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// 也可以创建自定义日志器在不同位置复用。
	// 新建日志器时可指定前缀，以便区分不同来源。
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// 现有 logger（包括标准 logger）也可以通过 `SetPrefix` 修改前缀（prefix）。
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// 日志器的输出目标可以自定义，只要实现 `io.Writer` 即可。
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// 这行日志会写入 `buf`。
	buflog.Println("hello")

	// 下行再将其打印到标准输出。
	fmt.Print("from buflog:", buf.String())

	// `slog` 包可输出结构化日志。
	// 例如使用 JSON 格式非常直接。
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// 除消息外，`slog` 还能附带任意数量的键值对。
	myslog.Info("hello again", "key", "val", "age", 25)
}
