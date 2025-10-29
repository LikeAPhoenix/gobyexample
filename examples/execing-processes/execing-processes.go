// 上一个示例介绍了[启动外部进程](spawning-processes)。
// 当需要外部进程与 Go 程序同时运行时会这么做。
// 有时我们想直接用另一个进程（可能不是 Go 实现）替换当前 Go 进程，
// 这时可以使用 Go 对经典 <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a> 函数的封装。

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// 示例中执行 `ls`。
	// Go 需要目标程序的绝对路径，因此使用 `exec.LookPath` 查找（通常是 `/bin/ls`）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` 接受切片形式的参数（而不是单个字符串）。
	// 给 `ls` 几个常见参数，注意第一个参数应为程序名。
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` 还需要[环境变量](environment-variables)。
	// 这里直接使用当前进程的环境。
	env := os.Environ()

	// 下面是真正的 `syscall.Exec` 调用。
	// 调用成功后，当前进程会被 `/bin/ls -a -l -h` 替换，不会再返回。
	// 若失败则会返回错误。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
