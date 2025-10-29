# 标准 Base64 与 URL Base64 输出略有差异（结尾 `+` vs `-`），但都能解码回原始字符串。
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
