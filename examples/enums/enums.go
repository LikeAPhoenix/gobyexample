// 枚举类型是[和类型](https://en.wikipedia.org/wiki/Algebraic_data_type)的特殊情况。
// 枚举的取值数量是固定的，每个值都有唯一名称。
// Go 没有单独的枚举语法，不过可以用现有惯用法轻松实现。

package main

import "fmt"

// 枚举类型 `ServerState` 的底层类型是 `int`。
type ServerState int

// 通过常量定义 `ServerState` 的所有可选值。
// 特殊关键字 [iota](https://go.dev/ref/spec#Iota) 会自动生成连续的常量值（本例为 0、1、2 等）。
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// 通过实现 [fmt.Stringer](https://pkg.go.dev/fmt#Stringer) 接口，
// 可以把 `ServerState` 的值打印或转换为字符串。
//
// 若状态很多，手动维护会比较繁琐，这时可以结合 `go:generate` 使用[stringer 工具](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)自动生成。
// 详情可参阅[这篇文章](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)。
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// 如果手里只有 `int` 类型的值，就无法传给 `transition`，编译器会报类型不匹配。
	// 这也提供了一定程度的编译期类型安全。

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// `transition` 模拟服务器状态转换，根据输入状态返回新状态。
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// 假设这里会检查一些条件以确定下一状态……
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
