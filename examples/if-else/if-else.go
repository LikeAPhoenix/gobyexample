// 在 Go 中使用 `if` 和 `else` 分支非常直接。

package main

import "fmt"

func main() {

	// 先看一个基础示例。
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// `if` 语句可以不带 `else`。
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 条件判断中常会用到逻辑运算符，比如 `&&` 和 `||`。
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// 在条件之前可以放置语句；在该语句中声明的变量在当前及之后的分支中都可见。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// 注意条件表达式外不需要括号，但 `{}` 是必不可少的。
