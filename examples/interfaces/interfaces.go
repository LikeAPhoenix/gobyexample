// 接口是方法签名的命名集合。

package main

import (
	"fmt"
	"math"
)

// 这里定义一个描述几何形状的基础接口。
type geometry interface {
	area() float64
	perim() float64
}

// 在示例中，我们为 `rect` 与 `circle` 类型实现该接口。
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 在 Go 中，只需实现接口中的所有方法即可视为实现该接口。
// 下面是 `rect` 对 `geometry` 的实现。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 若变量的类型为接口，就可以调用接口中声明的方法。
// `measure` 利用了这一点，可作用于任意 `geometry`。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// 有时需要获知接口值的运行时类型。
// 可以使用 type assertion（类型断言），或使用 type switch（类型 switch）。
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// `circle` 与 `rect` 类型都实现了 `geometry` 接口，
	// 因此它们的实例都可以传给 `measure`。
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
