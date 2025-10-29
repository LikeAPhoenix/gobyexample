$ go build command-line-subcommands.go 

# 先调用 foo 子命令。
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# 再试试 bar。
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# 注意 bar 不接受 foo 的标志。
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# 接下来介绍另一种常见的参数化方式：环境变量。
