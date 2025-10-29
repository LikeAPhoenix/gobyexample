// 在 Go 中写文件的模式与之前读取文件类似。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 首先展示如何把字符串（或字节）写入文件。
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 若需更细粒度的写入，可以先打开文件。
	f, err := os.Create("/tmp/dat2")
	check(err)

	// 惯例是在打开文件后立刻用 `defer` 安排关闭。
	defer f.Close()

	// 可以按预期使用 `Write` 写入字节 slice。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// 同样提供了 `WriteString` 方法。
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用 `Sync` 确保数据刷新到稳定存储。
	f.Sync()

	// `bufio` 除了提供缓冲读取器，也提供缓冲写入器。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// `Flush` 可以确保缓冲区的数据真正写入到底层 writer。
	w.Flush()

}
