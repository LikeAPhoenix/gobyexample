# 运行程序会计算哈希，并以易读的十六进制形式打印结果。
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# 其他哈希的用法与此类似。
# 例如要计算 SHA512，导入 `crypto/sha512` 并调用 `sha512.New()`。

# 若有密码学安全需求，请务必深入了解[哈希强度](https://en.wikipedia.org/wiki/Cryptographic_hash_function)。
