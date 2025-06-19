# 🧪 Go Beta: 学习 & 实验仓库

[![Go Version](https://img.shields.io/badge/go-1.21+-blue)](https://golang.org/doc/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Stars](https://img.shields.io/github/stars/ChenYujunjks/go_beta?style=social)](https://github.com/ChenYujunjks/go_beta/stargazers)

这是一个面向正在学习 Go 的开发者的实践型项目仓库，涵盖从基础语法、Web 框架（Gin）、并发编程，到 CLI 工具、JWT 认证、泛型等内容。  
本仓库展示了我在学习 Go 过程中的代码尝试与模块拆解，适合希望快速上手并参考实战项目结构的朋友。

> A personal Go learning & experimentation repo: exploring syntax, Gin, goroutines, CLI, generics, JWT, and more.

---

## 📁 项目结构 Project Structure

```

go-beta/
│
├── legacy/             # 初期结构：多个独立学习模块
│   ├── gointro/        # Go 入门语法与类型实践
│   ├── ginhello/       # 使用 Gin 编写基础 Web API
│   ├── cli/            # 自建 CLI 工具示例
│   ├── jwt/            # JWT 登录认证机制实现        
│   └── ui/      
│
├── generics/           # Go 1.18+ 泛型实践
├── goroutine/          # Goroutine 并发与 Channel 练习
├── defer/              # defer 执行机制与作用域分析
└── README.md           # 项目说明

````

---

## 🎯 学习目标 Learning Goals

- 熟悉 Go 的基本语法与核心特性（defer、struct、error 等）
- 掌握 Gin 框架构建 RESTful API
- 理解 Goroutine 并发模型与 Channel 通信
- 使用泛型提升代码复用与抽象能力
- 实现 JWT 登录认证的基本流程
- 了解 CLI 工具开发流程（flag、cobra 等）
- 通过项目实战逐步掌握真实开发技巧

---

## 💡 适合人群

* 正在学习 Go 编程语言的新手或自学者
* 想通过实战了解 Gin、JWT、Goroutine、泛型等 Go 特性的开发者
* 喜欢“边学边做”的编程方式

---

欢迎 Star 或 Fork，如你对某个模块感兴趣也欢迎提 issue 或交流建议 🙌


