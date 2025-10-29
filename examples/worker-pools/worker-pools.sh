# 运行结果显示 5 个任务被不同工人处理。
# 尽管总工作量约为 5 秒，但程序只花了 2 秒左右，
# 因为有 3 个工人并发执行。
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
