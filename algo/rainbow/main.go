package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var rainbowTable = make(map[string]string)

// 哈希函数
func hashMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

// 简单的还原函数，将哈希值转为明文（为了简化例子，仅返回哈希的前4位）
func reduceHash(hash string, index int) string {
	return hash[0:4] + fmt.Sprintf("%d", index)
}

// 生成彩虹表
func generateRainbowTable(passwords []string, chainLength int) {
	for _, password := range passwords {
		currentPassword := password
		var finalHash string
		for i := 0; i < chainLength; i++ {
			finalHash = hashMD5(currentPassword)
			currentPassword = reduceHash(finalHash, i)
		}
		// 存储起点和终点
		rainbowTable[finalHash] = password
	}
}

// 使用彩虹表查找哈希
func lookupHash(targetHash string, chainLength int) string {
	for i := chainLength - 1; i >= 0; i-- {
		currentHash := targetHash
		for j := i; j < chainLength; j++ {
			if password, found := rainbowTable[currentHash]; found {
				// 回溯链条找出明文
				for k := 0; k < j; k++ {
					currentHash = hashMD5(password)
					password = reduceHash(currentHash, k)
				}
				if hashMD5(password) == targetHash {
					return password
				}
			}
			currentHash = reduceHash(currentHash, j)
		}
	}
	return ""
}

func main() {
	// 用于生成彩虹表的明文密码
	passwords := []string{"password", "123456", "qwerty", "abc123"}
	chainLength := 5

	// 生成彩虹表
	generateRainbowTable(passwords, chainLength)

	// 目标哈希（假设是某个密码的哈希值）
	targetHash := hashMD5("password")
	fmt.Printf("目标哈希: %s\n", targetHash)

	// 查找哈希对应的明文密码
	result := lookupHash(targetHash, chainLength)
	if result != "" {
		fmt.Printf("找到匹配的明文密码: %s\n", result)
	} else {
		fmt.Println("未找到匹配的明文密码")
	}
}
