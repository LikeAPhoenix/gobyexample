// 在程序运行过程中，常会创建一些在退出后就无需保留的数据。
// 临时文件和目录非常适合这类需求，可以避免长久污染文件系统。

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 创建临时文件最简单的方式是调用 `os.CreateTemp`。
	// 它会创建文件并同时打开读写。
	// 传入 `""` 作为第一个参数，表示在操作系统默认位置创建。
	f, err := os.CreateTemp("", "sample")
	check(err)

	// 打印临时文件名。
	// 在类 Unix 系统中目录通常是 `/tmp`。
	// 文件名以第二个参数指定的前缀开头，其余部分自动生成，以确保并发调用产生不同名称。
	fmt.Println("Temp file name:", f.Name())

	// 使用结束后手动清理。
	// 操作系统可能会自动删除临时文件，但显式移除是更好的实践。
	defer os.Remove(f.Name())

	// 向文件写入数据。
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 如果需要写很多临时文件，可以先创建临时目录。
	// `os.MkdirTemp` 的参数与 `CreateTemp` 相同，但返回的是目录名。
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// 现在可以在临时目录下组合生成文件名。
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
