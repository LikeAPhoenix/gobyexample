# 翻译进度追踪
- 范围：仅翻译 `examples/` 目录内示例代码附带的注释与说明文字，以及 `templates/` 中的界面文案；`public/`、`README.md`、`CONTRIBUTING.md` 保持原样。
- 工具：直接编辑源文件，保留现有缩进和空行；必要时可使用 GoFmt 验证格式（不改变逻辑）。

## 术语与格式约定
- 注释、shell 脚本输出说明、模板中的可见文本翻译为中文；代码中的标识符与执行输出字符串保持英文，除非确为用户可读说明且不影响程序行为。
- Go 关键字与标准库函数不翻译；涉及专有名词、库函数、技术术语时优先保留英文原文，必要时可在中文括号中补充说明。
- 行内英文保留格式，中文标点使用全角，代码片段前后保持空格。
- 若需新增术语说明，请记录原文及说明方式，确保后续使用一致。

### 术语对照（统一风格）
- slice：统一用英文 “slice”；“slice 操作符”写作 `s[low:high]`。
- map：统一用英文 “map”。
- struct / interface：统一保留英文（中文语境可称“结构体/接口”，首次出现不强制括注）。
- goroutine / channel：统一保留英文。
- range：统一用英文 “range”（动词化表达仍写“使用 range 遍历”）。
- append：统一用英文 “append”，强调“需接收返回值”。
- rune：统一保留英文 “rune”（可描述为“Unicode 码点”）。
- UTF-8：统一保留英文。
- zero value：首次出现写“零值（zero value）”，其后使用“零值”。
- printf verb：写作“格式化动词（verb）”，具体动词用 `%d`、`%s` 等原样代码体例。
- Atoi / ParseXxx：函数名统一保留英文；必要时补充“十进制”“基数 0 自动推断”等中文说明。
- WaitGroup / Mutex / Context：统一保留英文；提及时可链接到对应示例。
- timer / ticker：统一保留英文；必要时在中文语境中补充“计时器/滴答器”。
- URL 术语：fragment、query（query string）等统一保留英文；必要时在中文中补充“片段/查询串”。
- sorting：Ordered（有序）的类型使用 `cmp.Ordered` 术语；`slices.SortFunc` 统一保留英文。
- random numbers：PCG 统一保留英文简称；`seed` 直接用英文。
- encoding/xml：字段 `XMLName`、标签 `id,attr` 等术语统一保留英文，并说明 attribute（属性）。
- JSON：`Marshal`/`Unmarshal`、`MarshalIndent` 统一保留英文；tag 使用反引号示例。
- Base64：统一保留英文 “Base64”；标准/URL 变体写作 `StdEncoding`/`URLEncoding`。
- SHA256：统一大写 “SHA256”；`crypto/sha256` 保留英文。
- text/template：action（动作）统一保留英文 action 并在首次括注；`Execute`、`Must` 保留英文。
- line filter：可译“行过滤器（line filter）”；stdin/stdout 建议英文附注。
- logging：logger（日志器）、flags、prefix 统一保留英文术语；结构化日志写作 structured logging。
- 并发（B4）
- goroutine / channel：统一保留英文；发送/接收写作“发送（send）/接收（receive）”。
- buffered/unbuffered channel：用“缓冲/无缓冲（buffered/unbuffered）channel”。
- channel 方向：send-only `chan<- T` / receive-only `<-chan T`。
- 关闭语义：closed channel、关闭后读取零值且 `ok == false`。
- range over channel：统一写法 “range 遍历 channel”。
- non-blocking：保留英文，可写“non-blocking 发送/接收”。
- select：统一保留英文 `select`/`case`/`default`。
- timeout：使用 `time.After`/`time.Timer` 描述超时（timeout）。
- timer / ticker：统一保留英文，必要时中文括注“计时器/滴答器”。
- rate limiting：rate limit/limiter、burst 统一保留英文。
- WaitGroup / Mutex / Atomic：统一保留英文。
- worker pool：可写“工作池（worker pool）”。
- 时间与上下文（B5）
- time / Duration：统一保留英文；描述持续时间时用 `time.Duration`，示例单位 ns/us/ms/s。
- layout：统一保留英文 layout（中文可写“布局”）；reference time 统一写作 `Mon Jan 2 15:04:05 MST 2006` 并括注 reference time。
- RFC3339：保留英文与常量名 `time.RFC3339`；UTC/本地（UTC/local）使用英文缩写。
- epoch：统一写作 Unix epoch；`Unix`/`UnixMilli`/`UnixNano` 保留英文。
- timer / ticker：统一保留英文；`NewTimer`/`NewTicker`、`Stop`/`Reset`、`C` 字段等 API 名称保留英文。
- timeouts：统一保留英文 timeout；`time.After`/`time.Timer` 用于实现超时。
- context：统一保留 `context.Context`；`Context()`、`Done()`、`Err()` 保留英文 API 名称。

- 系统与文件（B6）
- 目录/路径：统一使用 `os`、`path/filepath` 术语；`ReadDir`/`DirEntry`/`WalkDir`/`Chdir` 保留英文。
- 临时文件/目录：`CreateTemp`/`MkdirTemp` 保留英文；清理用 `defer` + `RemoveAll`。
- 文件读写：`os.Open`/`os.ReadFile`/`os.WriteFile`、`bufio.Scanner`/`Reader`/`Writer`、`io.Copy` 保留英文。
- 命令行：`os.Args`、`flag` 包（`Bool`/`Int`/`String`/`Parse`）、subcommand 统一英文。
- 环境变量：`os.Setenv`/`Getenv`/`Environ` 保留英文；键/值使用英文 key/value。
- embed：`//go:embed` 指令、`embed.FS` 保留英文。
- 进程：`exec.Command`、`Run`/`Start`/`Output`、`StdinPipe`/`StdoutPipe` 保留英文；子进程用 process。
- exec：`syscall.Exec`/`exec` 语义保持英文，“替换当前进程”用英文括注 replace current process。
- 退出：`os.Exit`、exit code（退出码）保留英文术语。
- 信号：`os/signal`、`Notify`、`SIGINT`/`SIGTERM` 保留英文；Ctrl+C 字面。
- HTTP 客户端：`http.Get`/`Post`/`Client`、`resp.Body.Close()` 保留英文。
- HTTP 服务端：`http.HandleFunc`/`ListenAndServe`、handler、router 保留英文；端口写 `:8090` 风格。
- URL：`net/url`、`Parse`/`RawQuery`/`SplitHostPort` 术语保持英文。
- 测试与基准：`*_test.go`、`go test`、`-bench`、`testing.B` 保留英文。
- variadic function：可译为“可变参数函数”，两种写法均可；标题或首次出现可写“variadic function（可变参数函数）”。
- receiver：方法接收者可写“接收者（receiver）”；文中多次出现时可直接使用“接收者”。
- type assertion / type switch：统一保留英文，可在首次出现时写“type assertion（类型断言）/ type switch（类型 switch）”。
- sentinel error：可写“哨兵错误（sentinel error）”。
- wrap error：统一使用“wrap（包装）错误”，可指明 `%w` in fmt.Errorf。
- iota：统一保留英文 iota；必要时括注“自动递增常量生成器”。
- Stringer：统一保留英文；必要时链接到 `fmt.Stringer` 文档。

- 泛型与迭代（B7）
- generics：首次出现可写“generics（类型参数 type parameters）”；`any`、`comparable`、`~T` 约束统一保留英文。
- type inference：统一用“type inference（类型推断）”。
- iterator：统一用英文 “iterator”；`iter.Seq`、`yield`、`slices.Collect` 保留英文。
- range：对内建类型与迭代器的 `range` 语义一致；rune/Unicode 保持英文术语。
- defer：统一保留英文 `defer`；与 `ensure`/`finally` 对照时用英文括注。

## 批次计划（examples/）
- [x] B1 基础语法与类型：`hello-world`、`values`、`variables`、`constants`、`for`、`if-else`、`switch`、`arrays`、`slices`、`maps`、`strings-and-runes`、`string-functions`、`string-formatting`、`number-parsing`
- [x] B2 函数与结构：`functions`、`multiple-return-values`、`variadic-functions`、`closures`、`recursion`、`pointers`、`methods`、`interfaces`、`structs`、`struct-embedding`、`enums`、`custom-errors`、`errors`、`panic`、`recover`
- [x] B3 数据处理：`sorting`、`sorting-by-functions`、`random-numbers`、`regular-expressions`、`json`、`xml`、`base64-encoding`、`sha256-hashes`、`text-templates`、`line-filters`、`logging`
- [x] B4 并发与同步：`goroutines`、`channels`、`channel-buffering`、`channel-directions`、`channel-synchronization`、`closing-channels`、`range-over-channels`、`non-blocking-channel-operations`、`select`、`timeouts`、`timers`、`tickers`、`rate-limiting`、`stateful-goroutines`、`atomic-counters`、`mutexes`、`waitgroups`、`worker-pools`
- [x] B5 时间与上下文：`time`、`time-formatting-parsing`、`epoch`、`timers`（如已完成标注）、`timeouts`（如已完成标注）、`context`、`rate-limiting`（如已完成标注）
- [x] B6 系统与文件：`directories`、`file-paths`、`reading-files`、`writing-files`、`temporary-files-and-directories`、`command-line-arguments`、`command-line-flags`、`command-line-subcommands`、`environment-variables`、`embed-directive`、`execing-processes`、`spawning-processes`、`exit`、`signals`、`http-client`、`http-server`、`url-parsing`、`testing-and-benchmarking`
- [x] B7 遗漏部分：`defer`、`generics`、`range-over-built-in-types`、`range-over-iterators`

> 注：批次间若目录重复，以首次所在批次为准，并在完成后于下方“进度备注”更新。

## 批次计划（templates/）
- [x] T1 站点主模板：`index.tmpl`、`example.tmpl`
- [x] T2 共享片段：`footer.tmpl`、`404.tmpl`
- [x] T3 资源检查：`site.js`、`site.css` 中可见文本（若有）

## 进度备注
- 2025-10-29：完成 B1 批次的注释与脚本翻译，并移除后续批次中的 `values` 重复项。
- 2025-10-29：完成 B2 批次（函数与结构相关目录）的翻译。
- 2025-10-29：完成 B3 批次（数据处理相关目录）的翻译。
- 2025-10-29：完成 B4 批次（并发与同步相关目录）的翻译。
- 2025-10-29：完成 B5 批次（时间与上下文相关目录）的翻译。
- 2025-10-29：完成 B6 批次（系统与文件相关目录）的翻译。
- 2025-10-29：完成 B7 批次（遗留条目）的翻译。
- 2025-10-29：完成 T1 批次（模板主页面）的翻译。
- 2025-10-29：完成 T2 批次（模板共享片段）翻译，并检查静态资源可见文本。
- 2025-10-29：复核术语及模板排版，`go test ./examples/...` 通过；`go test ./...` 因 tools 目录包含多个 `main` 函数无法构建，已记录。
