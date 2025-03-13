package main

import "fmt"

type Speaker interface {
	Speak() string
}

// Dog 类型实现了 Speaker 接口
type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪"
}

// Cat 类型也实现了 Speaker 接口
type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵"
}

// makeSpeak 函数接受一个 Speaker 接口类型的参数
func makeSpeak(s Speaker) {
	fmt.Println("动物说:", s.Speak())
}

func coupling() {
	// 声明一个 Speaker 类型的变量
	var s Speaker

	// 赋值为 Dog 类型的实例
	s = Dog{}
	fmt.Println(s.Speak()) // 输出：汪汪

	// 赋值为 Cat 类型的实例
	s = Cat{}
	fmt.Println(s.Speak()) // 输出：喵喵

	// 使用接口作为函数参数，实现多态调用
	makeSpeak(Dog{})
	makeSpeak(Cat{})
}
