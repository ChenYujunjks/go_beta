package main

import (
	"fmt"
	"os"

	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		//fatih color
		color.Red("这是红色的输出")
		color.Green("成功！")
		color.Cyan("信息：%s", "处理完成")
		fmt.Println("No command provided")
		//briandowns spinner
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // 可选多种样式
		s.Start()		
		time.Sleep(2 * time.Second) // 模拟长任务
		s.Stop()
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
