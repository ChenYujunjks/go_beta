package main

// 定义一个函数类型，接受字符串并返回字符串
type CommandFunc func(string) string

// 一个实现 CommandFunc 类型的函数
func sayHello(name string) string {
	return "Hello, " + name
}

// 使用自定义函数类型
var command CommandFunc = sayHello
