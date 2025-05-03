package main

import "fmt"

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
	printAnything(1)
	printAnything("hello")
	printAnything(true)
	printType(1)
	printType("hello")
	printType(true)
}
