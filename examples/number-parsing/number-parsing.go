// 在许多程序中，从字符串解析数字是常见的基础操作。
// 以下示例展示在 Go 中如何完成这件事。

package main

// 内建包 `strconv` 提供了相关的解析函数。
import (
	"fmt"
	"strconv"
)

func main() {

	// `ParseFloat` 的第二个参数 `64` 指定了解析的精度位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt` 的基数参数为 `0` 时，会根据字符串自行判断进制。
	// `64` 表示解析结果必须能放入 64 位。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` 也能识别十六进制格式的数字。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// 无符号整数可以使用 `ParseUint`。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` 是解析十进制 `int` 的便捷函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 当输入非法时，解析函数会返回错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
