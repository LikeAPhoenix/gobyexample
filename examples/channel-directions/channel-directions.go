// 将 channel 作为函数参数时，可以指定它只用于发送还是接收。
// 这种限定会提升程序的类型安全性。

package main

import "fmt"

// `ping` 只接受用于发送的 channel。
// 若尝试在该 channel 上接收，编译器会报错。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong` 函数则接收一个用于读取的 channel `pings`，以及一个用于发送的 channel `pongs`。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
