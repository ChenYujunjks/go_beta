package funcc

import (
	"errors"
	"strings"
)

func wordCount(s string) int {
	// 实现函数
	words := strings.Fields(s)
	return len(words)
}

func maxInArray(arr []int) (int, error) {
	var maxi int = 0
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}

	for _, value := range arr { //忽略索引，只使用值
		if value > maxi {
			maxi = value
		}
	}
	return maxi, nil
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	result := true
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			result = false
			break
		}
	}
	return result
}

func fibonacci(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
