// channel 默认是无缓冲（unbuffered）的，只有在有相应接收者准备好时才会接受发送。
// 而缓冲（buffered）channel 允许在没有立即接收者的情况下缓存有限数量的值。

package main

import "fmt"

func main() {

	// 这里创建一个最多可缓存 2 个字符串的 channel。
	messages := make(chan string, 2)

	// 由于是缓冲 channel，可以在没有同时接收者的情况下发送值。
	messages <- "buffered"
	messages <- "channel"

	// 随后再按常规方式接收即可。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
