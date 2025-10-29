// 标准库的 `strings` 包提供了许多实用的字符串处理函数。
// 下面通过几个例子快速了解它的功能。

package main

import (
	"fmt"
	s "strings"
)

// 下面的示例会频繁调用 `fmt.Println`，因此给它起个短别名。
var p = fmt.Println

func main() {

	// 下面展示 `strings` 包中的部分函数。
	// 这些都是包级函数，而非字符串类型的方法，所以需要把目标字符串作为第一个实参传入。
	// 更多函数可参阅 [`strings`](https://pkg.go.dev/strings) 包文档。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
