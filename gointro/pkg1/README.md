# Go Projects

## Introduction 简介

This repository contains various Go projects and learning materials. It includes a basic blog project written purely in Go, a project using the Gin framework, and several learning exercises and notes.

这个仓库包含了多个Go项目和学习资料。其中包括一个用纯Go语言编写的基本博客项目，一个使用Gin框架的项目，以及一些学习练习和笔记。

## Project Structure 项目结构

```plaintext
.
├── README.md
├── go_blog
│   ├── go.mod
│   ├── main.go
│   ├── public
│   │   └── resource
│   └── template
├── hello_gin
│   ├── fakemain.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── templates
│       ├── login.html
│       └── register.html
└── intro
    ├── day1
    ├── day2_ds
    ├── day3
    │   ├── main.go
    │   └── person.go
    ├── day4_error_interface
    │   ├── day4.md
    │   ├── error
    │   ├── interface.md
    │   └── main.go
    ├── day6
    ├── go.mod
    ├── go.sum
    └── notes
```

## Projects 项目

### 1. Go Blog 项目

This is a basic blog project implemented using only the core Go packages, without any additional frameworks like Gin.

这是一个使用纯Go语言包实现的基本博客项目，没有使用Gin等额外的框架。

- **Path 路径**: `./go_blog`
- **Description 描述**: Implements basic blog functionalities such as listing posts, viewing details, and user login.
- **主要功能**: 实现了博客的基本功能，如文章列表、详情查看和用户登录。

### 2. Hello Gin 项目

This project is an exploration of the Gin framework, demonstrating how to handle backend development with Gin.

这个项目是对Gin框架的探索，演示了如何使用Gin进行后端开发。

- **Path 路径**: `./hello_gin`
- **Description 描述**: Includes basic routes and templates for login and registration.
- **主要功能**: 包含了登录和注册的基本路由和模板。

### 3. Intro 项目

This directory contains my personal notes and code exercises while learning Go. It includes various exercises organized by day.

这个目录包含了我学习Go时的个人笔记和代码练习。它包括按天组织的各种练习。

- **Path 路径**: `./intro`
- **Description 描述**: Notes and exercises on Go concepts such as data structures, concurrency, interfaces, and more.
- **主要内容**: 关于Go语言概念的笔记和练习，如数据结构、并发、接口等。

## How to Run 运行方式

To run any of these projects, navigate to the project's directory and use the Go command:

要运行这些项目中的任何一个，请导航到项目目录并使用Go命令：

```sh
go run main.go
```

Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/).

确保你已在机器上安装Go。你可以从[golang.org](https://golang.org/)下载。

## Contributing 贡献

Feel free to fork this repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

欢迎fork这个仓库并提交pull request。如有重大更改，请先打开一个issue以讨论你想做的更改。

## License 许可证

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

这个项目是根据MIT许可证授权的——详细信息请参见[LICENSE](LICENSE)文件。