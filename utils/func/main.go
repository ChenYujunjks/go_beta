package funcc

import "fmt"

func stringLength(s string) int {
	return len(s)
}

func Day1_Func() {
	// 调用 prime.go 中的 isPrime 函数
	fmt.Println("29是否为素数:", isPrime(29))
	fmt.Println("151是否为素数:", isPrime(151))

	// 调用 fibonacci.go 中的 fibonacci 函数
	fmt.Println("Fibonacci(10) =", fibonacci(10))

	// 调用 max.go 中的 maxInArray 函数
	arr := []int{1, 5, 3, 9, 2}
	max, err := maxInArray(arr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("数组中的最大值：", max)
	}

	// 调用 string_length.go 中的 stringLength 函数
	str := "Hello, World!"
	fmt.Println("字符串长度：", stringLength(str))

	// 调用 word_count.go 中的 wordCount 函数
	sentence := "Go is an open-source programming language"
	fmt.Println("单词数量：", wordCount(sentence))
}
