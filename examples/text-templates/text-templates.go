// `text/template` 包提供生成动态内容或自定义输出的能力。
// 同名的 `html/template` 具有相同 API 并增加安全特性，适用于生成 HTML。

package main

import (
	"os"
	"text/template"
)

func main() {

	// 先创建模板并从字符串解析模板内容。
	// 模板由静态文本与 `{{...}}` 包裹的“动作”组成，用于动态插入内容。
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// 也可以使用 `template.Must`，在 `Parse` 出错时 panic。
	// 这对在全局作用域初始化模板特别方便。
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 执行模板时会将动作替换为指定值。
	// `{{.}}` 会被 `Execute` 的参数替代。
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// 辅助函数，便于后续创建模板。
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 若传入 struct，可通过 `{{.FieldName}}` 访问字段。
	// 要在模板中访问字段，struct 字段必须导出。
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// map 也适用同样的规则，对键名大小写没有限制。
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// `if/else` 实现条件判断。
	// 当值为类型的零值（如 0、空字符串、nil 指针等）时视为 false。
	// 此处还演示使用动作中的 `-` 去除空白。
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// `range` 可遍历 slice、数组、map 或 channel。
	// 在 `range` 块内，`{{.}}` 表示当前元素。
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
