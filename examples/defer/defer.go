// 延迟调用用于确保某个函数在稍后执行，通常用于清理资源。
// 在其他语言中常用的 `ensure` 或 `finally`，在 Go 中通常由 `defer` 承担。

package main

import (
	"fmt"
	"os"
)

// 假设我们需要先创建文件、写入内容，最后关闭文件。
// 可以借助 `defer` 轻松完成这一流程。
func main() {

	// 调用 `createFile` 获取文件对象后，立即通过 `defer closeFile` 延迟关闭。
	// 延迟调用会在当前函数（`main`）结束时执行，也就是 `writeFile` 完成之后。
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// 即便在延迟调用中，关闭文件时仍然应该检查错误。
	if err != nil {
		panic(err)
	}
}
