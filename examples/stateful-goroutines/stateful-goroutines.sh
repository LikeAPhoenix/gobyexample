# 运行结果显示，这个基于 goroutine 的状态管理示例共完成约八万次操作。
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177

# 在这个案例中，goroutine 方案比基于 mutex 的实现稍复杂一些。
# 但在需要与其他 channel 交互，或管理多把 mutex 易出错的场景，它会更有优势。
# 可以根据程序的可理解性与正确性需求选择合适的方案。
