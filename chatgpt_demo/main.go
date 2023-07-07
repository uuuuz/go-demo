package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	r := gin.Default()

	r.GET("/ai/query", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		message := ctx.Query("message")

		fmt.Println(message)
		res, err := server(message)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusBadRequest)
			_, _ = ctx.Writer.WriteString("please retry!")
			return
		}

		fmt.Println(res)
		ctx.JSON(http.StatusOK, map[string]string{"data": res})
	})

	_ = http.ListenAndServe(":10001", r)
}

func server(question string) (string, error) {
	//data, err := ioutil.ReadFile("chatgpt_demo/demo1.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//messages := []map[string]string{
	//	{"role": "user", "content": string(data)},
	//}
	//res, err := callChatGpt(messages)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("reply: ", res)

	messages := []map[string]string{
		//{"role": "system", "content": "以下是我提供的规范。根据这个规范回答问题。"},
		//{"role": "assistant", "content": res},
		{"role": "user", "content": question},
	}
	res, err := callChatGpt(messages)
	if err != nil {
		return "", err
	}
	return res, nil
}

var transport *http.Transport

func callChatGpt(messages []map[string]string) (string, error) {

	// 设置API访问参数
	//apiURL := "https://api.openai.com/v1/chat/completions"
	apiURL := "https://api.openai-sb.com/v1/chat/completions"
	//apiKey := "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"
	apiKey := "sb-29b8d019e623a4b6e4dbbbd0dff3d9b78cc8dffd7fab6248"

	// 构建请求数据
	data := map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": messages,
	}

	// 将数据编码为JSON
	jsonData, _ := json.Marshal(data)

	// 创建HTTP POST请求
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	if transport == nil {
		proxyURL, err := url.Parse("http://127.0.0.1:7890")
		if err != nil {
			return "", err
		}
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}
	client := http.Client{Transport: transport}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 读取响应内容
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	// 解析JSON响应
	var result map[string]interface{}
	_ = json.Unmarshal(body, &result)

	// 提取生成的回复
	choices := result["choices"].([]interface{})
	reply := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"]

	return fmt.Sprintf("%v", reply), nil
}
