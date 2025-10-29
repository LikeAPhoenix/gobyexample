# 为方便测试命令行标志，先编译再运行可执行文件。
$ go build command-line-flags.go

# 先为所有标志提供取值。
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# 如果省略某些标志，它们会使用默认值。
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# 标志之后还可以继续写位置参数。
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# `flag` 包要求所有标志都出现在位置参数之前，否则会被当成位置参数处理。
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# 使用 `-h` 或 `--help` 可查看自动生成的帮助信息。
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# 如果传入未定义的标志，程序会报错并再次显示帮助信息。
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
