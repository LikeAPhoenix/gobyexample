// [SHA256 哈希](https://en.wikipedia.org/wiki/SHA-2) 常用于为二进制或文本数据生成短标识。
// 例如 TLS/SSL 证书会使用 SHA256 计算签名。
// 下面演示如何在 Go 中计算 SHA256。

package main

// Go 在多个 `crypto/*` 包中提供各种哈希函数。
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// 从一个新的哈希实例开始。
	h := sha256.New()

	// `Write` 需要字节切片，若手头是字符串 `s`，请使用 `[]byte(s)`。
	h.Write([]byte(s))

	// `Sum` 会返回最终哈希结果的字节切片。
	// 它的参数可用于在现有切片后追加数据，这里传 `nil` 即可。
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
