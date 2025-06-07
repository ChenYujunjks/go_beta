# Go Beta 项目

这是一个用于学习和探索 Go 语言的项目仓库，涵盖了从基础语法到 Web 框架、CLI 工具、并发编程等多个方面的实践案例。

This is a personal Go learning and experimentation repo, including basic syntax, web frameworks, CLI tools, concurrency models, and more.

---

## 📁 项目结构 Project Structure

```
go-beta/
│
├── legacy/                 # Go 第一版代码，结构较旧
│   ├── algorithm/          # 算法练习与实现（排序、搜索等）
│   ├── ginhello/           # 使用 Gin Web 框架的示例
│   ├── gointro/            # Go 入门教程代码
│   ├── swagger/            # Swagger API 文档集成示例
│   ├── ui/                 # 与前端或终端交互相关的代码
│   └── cli/                # 命令行工具开发示例
│
├── defer/                  # Go 中 defer 用法的深入理解与实验
├── generics/               # Go 泛型功能的学习与应用（Go 1.18+）
├── jwt/                    # JWT 身份认证机制的实现与测试
├── goroutine/              # 并发编程：Goroutine 与 Channel 示例
└── README.md               # 项目说明文档
```

---

## 🧠 学习目标 Learning Goals

- 熟悉 Go 的基础语法和编程范式
- 掌握 Gin Web 框架开发 API
- 理解并发模型，实践 Goroutines 和 Channels
- 使用 Swagger 进行接口文档自动化
- 利用泛型提升代码复用性
- 实现 JWT 登录认证流程
- 构建基本 CLI 工具
- 掌握 defer 的底层机制和使用场景

---

## 🚀 快速开始 Quick Start

本项目没有统一的主程序，建议进入各个子文件夹阅读代码并逐个运行：

```bash
cd legacy/ginhello
go run main.go
```

如需安装依赖，可使用：

```bash
go mod tidy
```
