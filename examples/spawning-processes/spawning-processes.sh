# 启动的程序会输出与直接在命令行运行时相同的结果。
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

# `date` 没有 `-x` 标志，因此会打印错误并以非零状态退出。
command exited with rc = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
