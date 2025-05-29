package main

import (
	"cmp"
	"fmt"
)

// ① 泛型函数：最小值
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// ② 泛型类型：简单栈
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }
func (s *Stack[T]) Pop() T {
	n := len(s.data) - 1
	v := s.data[n]
	s.data = s.data[:n]
	return v
}

func main() {
	fmt.Println(Min(3, 7))          // int
	fmt.Println(Min(2.3, 1.8))      // float64

	var s Stack[string]
	s.Push("hi")
	s.Push("GPT")
	fmt.Println(s.Pop()) // GPT
}
