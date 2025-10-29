// Go 提供了许多用于操作文件系统目录的实用函数。

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 在当前工作目录下创建一个子目录。
	err := os.Mkdir("subdir", 0755)
	check(err)

	// 创建临时目录时，最好使用 `defer` 延迟删除。
	// `os.RemoveAll` 会删除整个目录树，类似于 `rm -rf`。
	defer os.RemoveAll("subdir")

	// 助手函数：创建空文件。
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// 使用 `MkdirAll` 可以一次创建多级目录（包含父级），类似命令行的 `mkdir -p`。
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` 会列出目录内容，返回 `os.DirEntry` slice。
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` 用于切换当前工作目录，等同于 `cd`。
	err = os.Chdir("subdir/parent/child")
	check(err)

	// 现在列出当前目录时，看到的是 `subdir/parent/child` 的内容。
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 切回初始目录。
	err = os.Chdir("../../..")
	check(err)

	// 还可以递归遍历目录及全部子目录。
	// `WalkDir` 接受回调，对每个文件或目录执行处理。
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// `visit` 会被 `filepath.WalkDir` 对每个遍历到的文件或目录调用。
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
