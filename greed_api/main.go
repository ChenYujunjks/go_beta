package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FearAndGreedIndex struct {
	Name string `json:"name"`
	Data []struct {
		Value               string `json:"value"`
		ValueClassification string `json:"value_classification"`
		Timestamp           string `json:"timestamp"`
		TimeUntilUpdate     string `json:"time_until_update,omitempty"`
	} `json:"data"`
	Metadata struct {
		Error interface{} `json:"error"`
	} `json:"metadata"`
}

func main() {
	// API URL
	url := "https://api.alternative.me/fng/?limit=1"

	// 创建 HTTP 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while making request:", err)
		return
	}
	defer resp.Body.Close()

	// 检查返回状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code")
		return
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response body:", err)
		return
	}

	// 解析 JSON 数据
	var index FearAndGreedIndex
	err = json.Unmarshal(body, &index)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 输出结果
	fmt.Printf("Name: %s\n", index.Name)
	for _, data := range index.Data {
		fmt.Printf("Value: %s, Classification: %s, Timestamp: %s\n", data.Value, data.ValueClassification, data.Timestamp)
	}
}
