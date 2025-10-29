// `switch` 语句可表达多分支条件逻辑。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 先看一个基础的 `switch`。
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在同一个 `case` 中可以用逗号分隔多个表达式。
	// 此外示例中还演示了可选的 `default` 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 省略表达式的 `switch` 是 `if/else` 的另一种写法。
	// 同时也展示了 `case` 中的表达式可以不是常量。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 类型 `switch` 比较的是类型而非值，可用来判断接口值的具体类型。
	// 在这个例子里，变量 `t` 会被推断为匹配分支对应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
