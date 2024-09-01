package utils

import (
	"fmt"
	"runtime"
)

func Anon_func() {
	m_maps := make(map[string]int)
	ch_m := make(chan int)
	fmt.Println(m_maps, ch_m)

	// 定义并立即执行一个匿名函数
	func() {
		fmt.Println("Hello, World!")
	}()

	// 带参数和返回值的匿名函数
	result := func(x, y int) int {
		return x + y
	}

	fmt.Println(result(3, 5)) // 输出

}

func Go_runtime() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPU)
}

func demoGarbageCollection() {
	s := make([]string, 0)
	for i := 0; i < 1000000; i++ {
		s = append(s, "This is a test string")
	}
	// 无需手动释放s，Go的垃圾回收器会处理未使用的内存
}

func defer_test() {
	fmt.Println("Start")
	defer fmt.Println("Deferred")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("End")
}
