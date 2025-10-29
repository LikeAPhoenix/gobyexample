// Go 的 struct 是带类型的字段集合，用于把相关数据组合成记录。

package main

import "fmt"

// `person` struct 包含 `name` 与 `age` 字段。
type person struct {
	name string
	age  int
}

// `newPerson` 根据提供的名字构造 `person` struct。
func newPerson(name string) *person {
	// Go 由垃圾回收器管理内存，因此可以安全地返回局部变量的指针。
	// 只要仍有引用，该变量就不会被清理。
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// 这种语法创建一个新的 struct 实例。
	fmt.Println(person{"Bob", 20})

	// 初始化时也可以显式指定字段名。
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段会填充零值。
	fmt.Println(person{name: "Fred"})

	// 使用 `&` 前缀可获得 struct 的指针。
	fmt.Println(&person{name: "Ann", age: 40})

	// 惯用做法是封装构造逻辑到工厂函数中。
	fmt.Println(newPerson("Jon"))

	// 通过点号访问 struct 字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 对 struct 指针同样使用点号，指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// struct 是可变的。
	sp.age = 51
	fmt.Println(sp.age)

	// 若 struct 类型只用于某个特定值，可以不给它命名，直接使用匿名 struct。
	// 这种技巧常用于[表驱动测试](testing-and-benchmarking)。
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
