package interf_test

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Speaker interface {
	Speak() string
}

// 实现接口的Speak方法
func (p Person) Speak() string {
	return "Hello my name is: " + p.Name
}

// 实现 **fmt.Stringer** 接口
func (p Person) String() string {
	return fmt.Sprintf("I realized the interface %s and I am %d years old", p.Name, p.Age)
}

func Greet(s Speaker) {
	fmt.Println(s.Speak())
}

func Intereiafoewj() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
	// 通过函数参数确定实现了接口
	Greet(p)

	var s Speaker = p
	// 使用类型断言
	s.Speak()
	if sp, ok := s.(Person); ok {
		fmt.Println("Person implements Speaker:", sp.Name, "年龄:", sp.Age)
	}
}
