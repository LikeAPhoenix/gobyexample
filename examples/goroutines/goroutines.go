// goroutine锛坓oroutine锛夋槸 Go 涓交閲忕骇鐨勬墽琛岀嚎绋嬨€?

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 鍋囪鎴戜滑鏈変竴涓嚱鏁拌皟鐢?`f(s)`锛屼互涓嬫槸鍚屾璋冪敤鐨勬柟寮忋€?
	f("direct")

	// 鑻ユ兂鍦?goroutine 涓皟鐢ㄥ嚱鏁帮紝鍙啓鎴?`go f(s)`銆?
	// 鏂板垱寤虹殑 goroutine 浼氫笌褰撳墠 goroutine 骞跺彂鎵ц銆?
	go f("goroutine")

	// 涔熷彲浠ュ鍖垮悕鍑芥暟鍚姩 goroutine銆?
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 鐜板湪涓や釜鍑芥暟璋冪敤閮藉湪鍚勮嚜鐨?goroutine 涓紓姝ヨ繍琛屻€?
	// 杩欓噷閫氳繃鐫＄湢绛夊緟瀹冧滑缁撴潫锛堟洿绋冲Ε鐨勫仛娉曟槸浣跨敤 [WaitGroup](waitgroups)锛夈€?
	time.Sleep(time.Second)
	fmt.Println("done")
}
