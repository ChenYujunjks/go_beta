package main

import (
	gor "beta/goroutine"
	utils "beta/pkg1"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("|----------------------------------|")
		fmt.Println("|请输入一个数字来选择要运行的模块: |")
		fmt.Println("|1: 普通                         |")
		fmt.Println("|2: Goroutine                   |")
		fmt.Println("|3: tBD      |")
		fmt.Println("|0: 退出                           |")
		fmt.Println("|----------------------------------|")
		// 获取用户输入
		fmt.Print("请输入选项: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // 去除空白字符

		// 尝试将输入转换为整数
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("请输入有效的数字！")
			continue // 跳过当前循环，继续要求用户输入
		}

		// 根据用户输入的数字来调用相应的函数
		switch choice {
		case 1:
			utils.Ask_one()
		case 2:
			gor.Ask_goroutine()
		case 3:
			fmt.Println("TBD!")
		case 0:
			fmt.Println("|--------------退出程序------------|")
			return
		default:
			fmt.Println("无效输入，请重新输入")
		}
	}
}
