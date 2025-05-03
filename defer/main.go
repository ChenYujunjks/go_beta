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

func main() {
	testDeferWithMutex()
	testDeferWithFile()
	testDeferWithLogging()
}
