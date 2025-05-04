package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var mu sync.Mutex

// 模拟锁的释放
func testDeferWithMutex() {
	fmt.Println("== 锁释放示例 ==")
	mu.Lock()
	fmt.Println("Locked")
	defer mu.Unlock() // 关键点：无论函数如何退出，都会释放锁
	fmt.Println("Doing something while locked")
}

// 模拟文件关闭
func testDeferWithFile() {
	fmt.Println("\n== 文件关闭示例 ==")
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保函数退出时关闭文件

	_, err = file.WriteString("Hello, Go!\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wrote to file")
}

func testDeferWithLogging() {
	fmt.Println("\n== 日志记录示例 ==")
	log.Println("Start processing")
	defer log.Println("Finished processing") // 结束时自动记录日志

	fmt.Println("Doing some work...")
}

func printAnything(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("int:", val)
	case string:
		fmt.Println("string:", val)
	default:
		fmt.Println("unknown type")
	}
}
func printType(v interface{}) {
	fmt.Printf("类型是: %T\n", v)
}
func main() {
	// 测试 printAnything 函数，传入不同类型的参数
	printAnything(1)       
	printAnything("hello")  
	printAnything(true)   

	// 测试 printType 函数，显示参数的具体类型
	printType(1)      
	printType("hello")    
	printType(true)  

	// 测试互斥锁的延迟释放
	testDeferWithMutex()   // 演示 defer 在锁释放中的应用

	// 测试文件的延迟关闭
	testDeferWithFile()    // 演示 defer 在文件操作中的应用

	// 测试日志的延迟记录
	testDeferWithLogging() // 演示 defer 在日志记录中的应用
}

