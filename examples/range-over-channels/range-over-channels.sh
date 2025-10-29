$ go run range-over-channels.go
one
two

# 本例也说明，即使在 channel 未读空时关闭它，剩余的值仍然可以被接收。
