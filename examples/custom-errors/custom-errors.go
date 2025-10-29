// 通过实现 `Error()` 方法，就能定义自定义错误类型。
// 本例在前面示例的基础上扩展，用自定义类型明确表示参数错误。

package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型通常以 “Error” 结尾。
type argError struct {
	arg     int
	message string
}

// 添加 `Error` 方法后，`argError` 就实现了 `error` 接口。
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// 返回自定义错误。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.As` 是 `errors.Is` 的进阶版本。
	// 它会检查某个错误（或其链上的任意错误）是否匹配指定类型，并尝试转换成该类型。
	// 匹配成功返回 `true`，否则返回 `false`。
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
