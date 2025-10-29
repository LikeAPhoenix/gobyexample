// 读写文件是许多 Go 程序的基础任务。
// 先来看读取文件的示例。

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件时大多需要检查错误。
// 这里提供一个辅助函数简化错误处理。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 最基础的读取方式是一次性把整个文件读入内存。
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 如果需要更精细的读取控制，可先通过 `Open` 获取 `os.File`。
	f, err := os.Open("/tmp/dat")
	check(err)

	// 从文件开头读取部分字节。
	// 这里最多读取 5 个，同时关注实际读取数量。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 可以使用 `Seek` 移动到指定位置再读取。
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// 也可以相对于当前游标移动，
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// 或相对于文件末尾移动。
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// `io` 包提供了更多实用函数。
	// 比如可以用 `ReadAtLeast` 更稳妥地实现类似的读取。
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 虽没有单独的“回到开头”函数，但 `Seek(0, io.SeekStart)` 即可实现。
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// `bufio` 包实现了带缓冲的读取器，既能提升大量小读操作的效率，也提供额外方法。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 操作结束后记得关闭文件（通常在 `Open` 后立即用 `defer` 调用）。
	f.Close()
}
