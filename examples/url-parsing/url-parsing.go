// URL 提供了[定位资源的统一方式](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)。
// 下面演示在 Go 中解析 URL。

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// 解析下例 URL，其中包含 scheme、认证信息、主机、端口、路径、查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析 URL，确认没有错误。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 获取 scheme 很直接。
	fmt.Println(u.Scheme)

	// `User` 包含认证信息，可调用 `Username` 和 `Password` 获取具体值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` 同时包含主机名和端口（若有）。
	// 使用 `SplitHostPort` 拆分。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 提取 `path` 以及 `#` 后的片段。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// `RawQuery` 返回 `k=v` 格式的查询串。
	// 也可解析为 map，键 map 到字符串 slice，若只需第一个值可取 `[0]`。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
