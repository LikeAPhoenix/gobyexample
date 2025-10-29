# 翻译进度追踪
- 范围：仅翻译 `examples/` 目录内示例代码附带的注释与说明文字，以及 `templates/` 中的界面文案；`public/`、`README.md`、`CONTRIBUTING.md` 保持原样。
- 工具：直接编辑源文件，保留现有缩进和空行；必要时可使用 GoFmt 验证格式（不改变逻辑）。

## 术语与格式约定
- 注释、shell 脚本输出说明、模板中的可见文本翻译为中文；代码中的标识符与执行输出字符串保持英文，除非确为用户可读说明且不影响程序行为。
- Go 关键字与标准库函数不翻译；涉及专有名词、库函数、技术术语时优先保留英文原文，必要时可在中文括号中补充说明。
- 行内英文保留格式，中文标点使用全角，代码片段前后保持空格。
- 若需新增术语说明，请记录原文及说明方式，确保后续使用一致。

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
