# 如预期一般，我们先收到 `"one"`，再收到 `"two"`。
$ time go run select.go 
received one
received two

# 注意总耗时约为 2 秒，因为两个 `Sleep` 是并发执行的。
real	0m2.245s
