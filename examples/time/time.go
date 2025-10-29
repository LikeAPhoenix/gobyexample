// Go 对时间与持续时间提供了丰富支持，下面是一些示例。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 先获取当前时间。
	now := time.Now()
	p(now)

	// 通过指定年、月、日等字段构造 `time`。
	// 时间总是与 `Location`（时区）关联。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 可按需提取时间的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 还可以获取星期信息（周一到周日）。
	p(then.Weekday())

	// 以下方法用于比较两个时间，判断先后或相等。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// `Sub` 返回一个 `Duration`，表示两个时间之间的间隔。
	diff := now.Sub(then)
	p(diff)

	// 可以以多种单位获取该时间间隔。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// `Add` 可以让时间前进或后退指定的持续时间。
	p(then.Add(diff))
	p(then.Add(-diff))
}
