// Go 支持在结构体类型上定义方法。

package main

import "fmt"

type rect struct {
	width, height int
}

// `area` 方法的接收者类型是 `*rect`。
func (r *rect) area() int {
	return r.width * r.height
}

// 方法既可以定义在指针接收者上，也可以定义在值接收者上。
// `perim` 就是值接收者的示例。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 调用刚才为结构体定义的两个方法。
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go 会在方法调用时自动处理值与指针之间的转换。
	// 如果希望避免调用时的拷贝，或让方法修改接收者，可以选择指针接收者。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
