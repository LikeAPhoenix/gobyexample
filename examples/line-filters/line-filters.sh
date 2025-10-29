# 想要体验这个行过滤器，先准备一个包含若干小写文本的文件。
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# 然后通过行过滤器将其转换为大写。
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
