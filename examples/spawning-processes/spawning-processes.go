// 有时 Go 程序需要启动其他进程。

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// 先从一个简单命令开始，不带参数也无需输入，仅向标准输出打印信息。
	// `exec.Command` 会返回表示该外部进程的对象。
	dateCmd := exec.Command("date")

	// `Output` 方法会运行命令、等待其结束并收集标准输出。
	// 若没有错误，`dateOut` 将包含日期信息的字节。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` 等方法在执行命令失败（例如路径不正确）时返回 `*exec.Error`；
	// 如果命令运行后以非零返回码退出，则返回 `*exec.ExitError`。
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		var execErr *exec.Error
		var exitErr *exec.ExitError
		switch {
		case errors.As(err, &execErr):
			fmt.Println("failed executing:", err)
		case errors.As(err, &exitErr):
			exitCode := exitErr.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		default:
			panic(err)
		}
	}

	// 接下来演示更复杂的情况：向外部进程的 `stdin` 写入数据，并从 `stdout` 收集结果。
	grepCmd := exec.Command("grep", "hello")

	// 显式获取输入/输出管道，启动进程，写入数据，读取输出，最后等待进程退出。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// 上述示例省略了错误处理，实际使用时仍可采用 `if err != nil` 模式。
	// 此处只收集了 `StdoutPipe`，同样方法也可用于 `StderrPipe`。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 注意启动命令时需要显式传入命令及参数切片，而非单个字符串。
	// 如果希望通过字符串执行完整命令，可以借助 `bash -c`：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
