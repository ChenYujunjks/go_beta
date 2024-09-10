package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		return
	}

	command := os.Args[1]

	switch command {
	case "im":
		printIm()
	default:
		fmt.Println("Unknown command")
	}
}

func printIm() {
	fmt.Println("Hello, this is the 'im' command")
}
