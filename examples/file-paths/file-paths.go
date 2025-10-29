// `filepath` 包提供了解析与构建文件路径的函数，具备跨操作系统的可移植性。
// 例如在 Linux 上是 `dir/file`，而 Windows 上是 `dir\file`。
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// 使用 `Join` 以可移植方式构建路径。
	// 它接受任意数量的参数并拼接成层级路径。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// 应始终使用 `Join`，而非手动拼接 `/` 或 `\`。
	// 除了可移植，`Join` 还会通过移除多余分隔符、目录变更来规范路径。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` 与 `Base` 可将路径拆分为目录与文件。
	// 或者使用 `Split` 一次返回两者。
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// 判断路径是否为绝对路径。
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// 某些文件名带扩展名，可通过 `Ext` 拆分出来。
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 若要去除扩展名，可使用 `strings.TrimSuffix`。
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` 用于求解基准路径与目标路径之间的相对路径。
	// 若无法表示相对关系则返回错误。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
