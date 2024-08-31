package goroutine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Ask_goroutine() {
	fmt.Println("|----------------------------------|")
	fmt.Println("|请输入一个数字来选择要运行的函数: |")
	fmt.Println("|1: 运行匿名函数                   |")
	fmt.Println("|2: 运行 runtime 查看CPU核心       |")
	fmt.Println("|3: Test MySql Connection          |")
	fmt.Println("|4: Test PostgreSql Connection     |")
	fmt.Println("|0: 退出                           |")
	fmt.Println("|----------------------------------|")

	// 获取用户输入
	reader := bufio.NewReader(os.Stdin)
	for {
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

		case 0:
			fmt.Println("|--------------退出程序------------|")
			return
		default:
			fmt.Println("无效输入，请重新输入")
		}
	}
}
