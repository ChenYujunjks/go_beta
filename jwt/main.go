package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("my_super_secret_key") // 就是 .env 里 JWT_SECRET 的内容

func main() {
	// 1. 签发 token
	tokenString, err := generateJWT("123")
	if err != nil {
		panic(err)
	}
	fmt.Println("🔑 Token:", tokenString)

	// 2. 验证 token
	userID, err := verifyJWT(tokenString)
	if err != nil {
		fmt.Println("❌ Invalid token:", err)
	} else {
		fmt.Println("✅ Token valid, userID =", userID)
	}
}

// 签发 JWT（包含 userID 作为 Subject）
func generateJWT(userID string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(), // 1小时过期
		IssuedAt:  time.Now().Unix(),
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 验证 JWT，返回 userID（从 claims.Subject 提取）
func verifyJWT(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil // 返回密钥用于验签
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims.Subject, nil
	} else {
		return "", fmt.Errorf("invalid token claims")
	}
}
