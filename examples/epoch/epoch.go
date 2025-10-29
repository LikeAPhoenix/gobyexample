// 程序中常需要获取自 [Unix 纪元（Unix epoch）](https://en.wikipedia.org/wiki/Unix_time) 以来的秒、毫秒或纳秒数。
// 下面演示在 Go 中如何实现。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 使用 `time.Now` 配合 `Unix`、`UnixMilli` 或 `UnixNano`，分别得到距 Unix 纪元的秒、毫秒与纳秒。
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// 也可以把自纪元以来的秒或纳秒转换回 `time` 类型。
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
