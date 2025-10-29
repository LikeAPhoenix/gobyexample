$ go run channel-synchronization.go      
working...done                  

# 如果移除程序中的 `<- done`，主程序可能会在 `worker` 完成甚至启动之前就退出。
