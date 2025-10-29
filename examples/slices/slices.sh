# 注意 slice 虽然不同于数组，但使用 `fmt.Println` 输出时形式相似。
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# 想了解 slice 的设计与实现细节，可阅读 Go 团队的[精彩博客文章](https://go.dev/blog/slices-intro)。

# 学习完数组和 slice 之后，接下来要看看 Go 的另一种核心内建数据结构：map。
