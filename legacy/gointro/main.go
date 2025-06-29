package main

import (
	er "beta/error"
	func_intro "beta/func"
	interf_test "beta/interface"
	struct_l "beta/struct"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("|您现在位于模块1：简单学习---------|")
	fmt.Println("|请输入一个数字来选择要运行的函数: |")
	fmt.Println("|1: 运行匿名函数                   |")
	fmt.Println("|2: 运行 runtime 查看CPU核心       |")
	fmt.Println("|3: 查看Day1 func                  |")
	fmt.Println("|4: 查看Day2 error                 |")
	fmt.Println("|5: 查看Day3 interface             |")
	fmt.Println("|6: 查看Day4 struct                |")
	fmt.Println("|7: 查看defer的用法                 |")
	fmt.Println("|0: 退出                           |")
	fmt.Println("|----------------------------------|")

	// 获取用户输入
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("---------------------------------->>")
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
			Anon_func()
		case 2:
			Go_runtime()
		case 3:
			func_intro.Day1_Func()
		case 4:
			er.Error_F()
		case 5:
			interf_test.Intereiafoewj()
		case 6:
			struct_l.Struct_jjj()
		case 7:
			defer_test()
		case 0:
			fmt.Println("|--------------退出程序------------|")
			return
		default:
			fmt.Println("无效输入，请重新输入")
		}
	}

}
