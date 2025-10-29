// Go 通过基于模式的布局支持时间格式化与解析。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 先根据 RFC3339 布局常量格式化当前时间。
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 时间解析与 `Format` 使用相同的布局值。
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` 与 `Parse` 使用示例驱动的布局。
	// 通常会用 `time` 包内置常量，但也可以提供自定义布局。
	// 布局必须基于参考时间 `Mon Jan 2 15:04:05 MST 2006`，以指示格式模式。
	// 该示例时间中的年份、小时、星期等必须按示例准确填入。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 如果需要纯数字表示，也可以取出各组件后使用常规字符串格式化。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` 在输入格式不正确时会返回错误，说明解析失败的原因。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
