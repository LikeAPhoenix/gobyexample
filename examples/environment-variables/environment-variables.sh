# 运行程序可看到在代码中设置的 `FOO` 已生效，而未设置的 `BAR` 为空。
$ go run environment-variables.go
FOO: 1
BAR: 

# 列出的环境变量键取决于当前机器。
TERM_PROGRAM
PATH
SHELL
...
FOO

# 如果事先在环境中设置 `BAR`，程序就能读到对应的值。
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
