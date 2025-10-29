// 有些命令行工具（如 `go` 或 `git`）拥有多个子命令，每个子命令都有自己的标志。
// 例如 `go build` 与 `go get` 是 `go` 工具的两个子命令。
// `flag` 包可以方便地定义带独立标志的子命令。

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 使用 `NewFlagSet` 声明子命令，并为其定义专属标志。
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// 其他子命令可以定义不同的标志。
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// 子命令应作为程序的第一个参数。
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// 根据第一个参数判断调用的是哪个子命令。
	switch os.Args[1] {

	// 每个子命令都解析自己的标志，并可获取剩余的位置参数。
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
