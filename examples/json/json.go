// Go 内建支持 JSON 的编码与解码，既适用于内建类型也适用于自定义类型。

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 下面通过两个 struct 演示自定义类型的编解码。
type response1 struct {
	Page   int
	Fruits []string
}

// 只有导出的字段才会参与 JSON 编解码。
// 要导出字段，名称必须以大写字母开头。
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// 先来看如何把基础数据类型编码成 JSON 字符串。
	// 以下是一些原子类型的示例。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// slice 和 map 会如预期那样编码为 JSON 数组和对象。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 包也能自动编码自定义类型。
	// 它只会包含导出的字段，并默认使用字段名作为键。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 在 struct 字段上使用标签可以自定义 JSON 键名。
	// 参见上方 `response2` 的定义。
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 接着将 JSON 数据解码为 Go 值。
	// 下面是一个解码到通用数据结构的例子。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 需要提供一个变量让 JSON 包写入解码结果。
	// 这里的 `map[string]interface{}` 会保存任意类型的数据。
	var dat map[string]interface{}

	// 下面执行解码，并检查是否出错。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 若想使用解码后的值，需要将其转换为合适的类型。
	// 比如这里把 `num` 转换为预期的 `float64`。
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套数据时则需要多次转换。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// JSON 也可以直接解码到自定义类型。
	// 这样做能提供额外的类型安全，并免去取值时的类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 上面的示例使用字节或字符串作为数据与 JSON 表示之间的中间层。
	// 也可以直接将 JSON 输出流式写入 `os.Writer`，例如 `os.Stdout` 或 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// 同理，可借助 `json.Decoder` 从 `os.Reader`（如 `os.Stdin` 或 HTTP 请求体）流式读取。
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
