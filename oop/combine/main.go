package main

import "fmt"

// 定义接口
type Animal interface {
	Speak() string
}

// 定义具体类型
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

// Cat 实现 Animal
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

// 函数只依赖于接口
func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	d := Dog{}
	c := Cat{}
	MakeAnimalSpeak(d) // 输出 "Woof"
	MakeAnimalSpeak(c) // 输出 "Meow"
}
