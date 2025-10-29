# `zeroval` 不会改变 `main` 中的 `i`，
# 而 `zeroptr` 会，因为它持有该变量的内存地址引用。
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
