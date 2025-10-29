// `encoding/xml` 包为 XML 及类似格式提供内建支持。

package main

import (
	"encoding/xml"
	"fmt"
)

// `Plant` 结构会映射到 XML。
// 与 JSON 示例类似，字段标签用于指导编解码器。
// 这里展示 XML 包的几个特性：字段名 `XMLName` 指定对应的 XML 元素名称；
// `id,attr` 表示 `Id` 字段应作为 XML 属性而非子元素。
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// 将结构体编码为 XML。
	// 使用 `MarshalIndent` 生成更易读的格式。
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// 若想添加通用的 XML 头部，可手动拼接。
	fmt.Println(xml.Header + string(out))

	// 使用 `Unmarshal` 将 XML 字节流解析到数据结构。
	// 如果 XML 不合法或无法映射到 `Plant`，会返回具体错误。
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// 标签 `parent>child>plant` 指示编码器将所有 `plant` 节点嵌套在 `<parent><child>...` 中。
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
