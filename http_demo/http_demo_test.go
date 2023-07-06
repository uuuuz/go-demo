package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func Test_http(t *testing.T) {
	body := &bytes.Buffer{}
	//write := multipart.NewWriter(body)
	//if err := write.WriteField("username", "wxm"); err != nil {
	//	log.Warn("err:", err)
	//	gtx.Status(500)
	//	return
	//}
	//if err := write.WriteField("password", "wxm"); err != nil {
	//	log.Warn("err:", err)
	//	gtx.Status(500)
	//	return
	//}
	//if err := write.Close(); err != nil {
	//	log.Warn("err:", err)
	//	gtx.Status(500)
	//	return
	//}

	val := url.Values{}
	//val.Set("username", "wxm")
	//val.Set("password", "wxm")
	body.WriteString(val.Encode())

	// 1
	resp, err := http.Post("http://localhost:8088/login/", "application/x-www-form-urlencoded", body)

	// 2
	//req, err := http.NewRequest("POST", "http://localhost:8088/login/", body)
	//if err != nil {
	//	t.Log(fmt.Sprintf("err: %v", err))
	//	return
	//}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	////req.AddCookie(&http.Cookie{
	////	Name:  "session",
	////	Value: "eyJsb2NhbGUiOiJlbiJ9.Y_XU4Q.di2HMI4Th8tIIfSOqWNSFQEFeJ4",
	////})
	//resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Log(fmt.Sprintf("err: %v", err))
		return
	}

	value := resp.Header.Get("Set-Cookie")
	t.Log(fmt.Sprintf("value: %s", value))

	//gtx.Request.Header.Set("Cookie", "session=.eJwlzjEKwzAMBdC7eO4gR7Zl5zLBkr5oIbSQ0Kn07g10fNv7pC0OnPe0xtxP3NL28LSmEDaI8VLUI7fZg8lpijLl6s6LWJYhlUuomnHrTa1n7RQLeFRURcClmhIPRjEYaCjNKLZwJyltTscYOr1mJjVtyNwZNSS3dEXeJ47_hsvl_WVzxyU80_cHUr41Lg.Y_SahQ.8cOq_LV51sjlFVR83hrvJr9AK5Y")
	//gtx.Redirect(http.StatusMovedPermanently, "http://localhost:8088/superset/dashboard/1/?standalone=true")
	return
}

func Test2(t *testing.T) {
	body := &bytes.Buffer{}
	//body.WriteString(val.Encode())
	body.Write([]byte(`{"username": "wxm","password": "wxm","provider": "db","refresh": false}`))

	// 1
	//resp, err := http.Post("http://localhost:8088/login/", "application/x-www-form-urlencoded", body)

	// 2
	req, err := http.NewRequest("POST", "http://localhost:8088/api/v1/security/login", body)
	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")
	//v, err := getSession()
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("err: %v", err))
	//	return
	//}
	//req.Header.Set("Cookie", strings.Split(v, ";")[0])
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return
	}
	var resMap = make(map[string]string)
	err = json.Unmarshal(data, &resMap)
	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return
	}
	t.Log(resMap)
}

func Test3(t *testing.T) {
	//body := &bytes.Buffer{}
	//write := multipart.NewWriter(body)
	//if err := write.WriteField("username", "wxm"); err != nil {
	//	fmt.Println(fmt.Sprintf("err: %v", err))
	//	return
	//}
	//if err := write.WriteField("password", "wxm"); err != nil {
	//	fmt.Println(fmt.Sprintf("err: %v", err))
	//	return
	//}
	//
	//authorizationToken, err := getAuthorizationToken()
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("err: %v", err))
	//	return
	//}
	//authorizationToken = "Bearer " + authorizationToken
	//
	//csrfToken, err := getCsrfToken(authorizationToken)
	//fmt.Println("csrfToken:" + csrfToken)
	//if err := write.WriteField("csrf_token", csrfToken); err != nil {
	//	fmt.Println(fmt.Sprintf("err: %v", err))
	//	return
	//}
	//if err := write.Close(); err != nil {
	//	fmt.Println(fmt.Sprintf("err: %v", err))
	//	return
	//}

	val := url.Values{}
	val.Set("username", "wxm")
	val.Set("password", "wxm")
	val.Set("csrf_token", "IjY3M2YxZDQyMGFjMDVkYWNjNmExNjQzYTQ0MmE3YTc0OTJlN2UwZDEi.Y_YaDQ.vAnlZj2xPlNvrs9XEzSfv_jN8Fs")
	body := &bytes.Buffer{}
	body.WriteString(val.Encode())

	// 1
	//resp, err := http.Post("http://localhost:8088/login/", "application/x-www-form-urlencoded", body)

	// 2
	req, err := http.NewRequest("POST", "http://localhost:8088/login/", body)
	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "multipart/form-data")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return
	}
	t.Log(resp)
}

// GOPROXY=https://goproxy.cn,direct
func Test_t(t *testing.T) {
	//authorizationToken, err := getAuthorizationToken()
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("err: %v", err))
	//	return
	//}

	// http://localhost:8088/api/v1/security/csrf_token
	// http://localhost:8088/api/v1/security/csrf_token
	req, err := http.NewRequest("GET", "http://10.110.1.86:8088/api/v1/security/csrf_token", nil)
	if err != nil {
		t.Log(err)
		return
	}
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6dHJ1ZSwiaWF0IjoxNjc3MTIyMzYxLCJqdGkiOiIwNjkzNzhjZi0yNzI2LTRiYzktYTU4Ny0xOTc4YWE5YWFiZTEiLCJ0eXBlIjoiYWNjZXNzIiwic3ViIjoxLCJuYmYiOjE2NzcxMjIzNjEsImV4cCI6MTY3NzEyMzI2MX0.ECDXW_Lrke7caTyCtBbTmMbFWM4U89Z1hHt2NgmHV3k")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(data))
	var result = make(map[string]string)
	if err = json.Unmarshal(data, &result); err != nil {

	}
	csrfToken := result["result"]

	//csrfToken, _ := getCsrfToken(authorizationToken)
	t.Log(csrfToken)
}

func Test4(t *testing.T) {
	authorizationToken, err := getAuthorizationToken()
	if err != nil {
		fmt.Println(fmt.Sprintf("err: %v", err))
		return
	}
	t.Log(authorizationToken)
}
