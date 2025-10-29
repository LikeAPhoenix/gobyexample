// `//go:embed` 是一种[编译器指令](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives)，
// 允许在构建时将任意文件或文件夹打包进 Go 可执行文件。
// 可在[此处](https://pkg.go.dev/embed)了解更多信息。
package main

// 引入 `embed` 包；若未使用其中导出的标识符，可使用 `_ "embed"` 的空导入。
import (
	"embed"
)

// `embed` 指令使用相对 Go 源文件所在目录的路径。
// 该指令会把对应文件内容嵌入紧随其后的 `string` 变量。
//
//go:embed folder/single_file.txt
var fileString string

// 也可以嵌入到 `[]byte` 变量中。
//
//go:embed folder/single_file.txt
var fileByte []byte

// 还可以通过通配符嵌入多个文件甚至整个文件夹。
// 这里使用 [embed.FS](https://pkg.go.dev/embed#FS) 类型，提供简单的虚拟文件系统。
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// 打印 `single_file.txt` 的内容。
	print(fileString)
	print(string(fileByte))

	// 从嵌入的目录中读取其他文件。
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
