// Go 惯用通过单独的返回值来传递错误。
// 这与 Java、Python、Ruby 等语言中的异常或 C 语言里复合返回值的做法不同。
// Go 的方式让我们一眼就能看出哪些函数会返回错误，并使用相同的语言结构处理错误与其他逻辑。
//
// 想了解更多细节，可参阅 [errors 包文档](https://pkg.go.dev/errors)以及[这篇博文](https://go.dev/blog/go1.13-errors)。

package main

import (
	"errors"
	"fmt"
)

// 约定上，错误是函数的最后一个返回值，类型为内建接口 `error`。
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` 会根据给定消息构造一个基础错误值。
		return -1, errors.New("can't work with 42")
	}

	// 当错误位置返回 `nil` 时表示没有发生错误。
	return arg + 3, nil
}

// 哨兵错误（sentinel error）是预先声明的变量，用于表示特定的错误情形。
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// 可以把底层错误包装在更高层的错误里，以补充上下文。
		// 最简单的方式是在 `fmt.Errorf` 中使用 `%w` 动词。
		// 被包装的错误形成一条链，可用 `errors.Is`、`errors.As` 查询。
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// 惯用写法是在 `if` 语句行内就完成错误检查。
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` 会检查错误（或其链上的任何错误）是否与指定错误值匹配。
			// 对于嵌套或包装的错误非常实用，可识别特定类型或哨兵错误。
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
