package main

import "fmt"

// 定义接口：能飞的能力
type Flyer interface {
	Fly() string
}

type Bird struct {
	Species string
}

func (b Bird) Fly() string {
	return b.Species + " is flying!"
}

type Airplane struct {
	Model string
}

func (a Airplane) Fly() string {
	return "Airplane " + a.Model + " is flying!"
}

// ✅ 示例1：让任意 Flyer 飞
func letItFly(f Flyer) {
	fmt.Println(f.Fly())
}

// ✅ 示例2：同时处理多个 Flyer
func flyAll(f []Flyer) {
	for _, f := range f {
		letItFly(f)
	}
}

func main() {
	// 示例1：多态函数调用
	fmt.Println("== 多态函数调用 letItFly 示例 ==")
	letItFly(Bird{Species: "Swan"})
	letItFly(Airplane{Model: "Airbus A320"})

	// 示例2：接口数组统一处理
	fmt.Println("\n== 接口数组统一处理 flyAll 示例 ==")
	flyers := []Flyer{
		Bird{Species: "Eagle"},
		Airplane{Model: "F-22"},
		Bird{Species: "Parrot"},
	}
	flyAll(flyers)
}
