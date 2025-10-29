// Go 内建支持[Base64 编码与解码](https://en.wikipedia.org/wiki/Base64)。

package main

// 这里将 `encoding/base64` 以 `b64` 别名导入，便于后续代码书写。
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// 这是待编码/解码的字符串。
	data := "abc123!?$*&()'-=@~"

	// Go 同时支持标准和 URL 兼容的 Base64。
	// 下面使用标准编码器，调用前需将字符串转换为 `[]byte`。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能返回错误；若无法确定输入是否合法，记得进行检查。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 接下来使用 URL 兼容的 Base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
