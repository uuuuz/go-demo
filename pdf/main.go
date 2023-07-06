package main

import (
	"bytes"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"os"
)

// 为pdf文件添加密码
// https://github.com/pdfcpu/pdfcpu
func main() {
	// 打开pdf文件
	f1, err := os.Open("pdf/test.pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f1.Close()
	// 加密 f1 并将结果输入到 buf
	buf := &bytes.Buffer{}
	// 使用默认配置，并设置密码
	cfg := api.LoadConfiguration()
	cfg.OwnerPW = "!@_/123*&#+A天天向上qs0-1sfk[]~df" // 拥有者密码
	cfg.UserPW = "!@_/123*&#+A天天向上qs0-1sfk[]~dd"  // 用户密码

	ori := bytes.NewReader([]byte{})
	err = api.Encrypt(ori, buf, cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 创建新文件，写入buf数据
	f2, err := os.Create("pdf/test_enc.pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f2.Close()
	_, err = f2.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
