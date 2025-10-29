// 单元测试是编写健壮 Go 程序的重要环节。
// `testing` 包提供了编写测试所需的工具，而 `go test` 命令负责执行。

// 为了演示，这段代码放在 `main` 包中，实际可以是任意包。
// 测试通常与被测代码同包。
package main

import (
	"fmt"
	"testing"
)

// 我们将测试一个求整型最小值的简单实现。
// 实际项目中，被测代码可能位于 `intutils.go`，
// 对应的测试文件命名为 `intutils_test.go`。
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 测试函数以 `Test` 前缀命名。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` 会报告失败但继续执行，`t.Fatal*` 会立即终止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 为避免重复，惯用“表驱动”风格：把输入与期望输出列成表，用循环执行测试逻辑。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// `t.Run` 可为每个表项运行子测试，执行 `go test -v` 时会单独展示。
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// 基准测试通常也放在 `_test.go` 文件中，名称以 `Benchmark` 开头。
// 运行基准前需要执行但不计时的准备工作应放在循环外。
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// 基准测试框架会自动多次执行循环体，以估算单次迭代的运行时间。
		IntMin(1, 2)
	}
}
