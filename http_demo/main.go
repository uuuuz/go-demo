package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var host = "10.110.1.86:8088"
var username = "admin"
var password = "admin"

func main() {
	r := gin.Default()
	r.GET("/test35", func(gtx *gin.Context) {

		authorizationToken, err := getAuthorizationToken()
		if err != nil {
			fmt.Println(fmt.Sprintf("err: %v", err))
			return
		}
		authorizationToken = "Bearer " + authorizationToken

		csrfToken, cookie, err := getCsrfToken(authorizationToken)
		fmt.Println(csrfToken)
		val := url.Values{}
		val.Set("username", username)
		val.Set("password", password)
		val.Set("csrf_token", csrfToken)
		body := &bytes.Buffer{}
		body.WriteString(val.Encode())

		req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/login/", host), body)
		if err != nil {
			fmt.Println(fmt.Sprintf("err: %v", err))
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", cookie)
		resp, err := getClientWithRedirect().Do(req)

		if err != nil {
			fmt.Println(fmt.Sprintf("err: %v", err))
			return
		}
		cookie = strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]
		fmt.Println(cookie)
		fmt.Println("cookie:" + cookie)
		fmt.Println(fmt.Sprintf("url: http://%s/superset/dashboard/1/?standalone=true", host))
	})
	_ = http.ListenAndServe(":8000", r)
}

func getAuthorizationToken() (string, error) {
	body := &bytes.Buffer{}
	body.Write([]byte(`{"username": "admin","password": "admin","provider": "db","refresh": false}`))

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/api/v1/security/login", host), body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := getClient().Do(req)
	//resp, err := http.Post(fmt.Sprintf("http://%s/api/v1/security/login", host), "application/json", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result = make(map[string]string)
	if err = json.Unmarshal(data, &result); err != nil {
		return "", err
	}
	return result["access_token"], nil
}

func getCsrfToken(token string) (string, string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/v1/security/csrf_token", host), nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Authorization", token)
	resp, err := getClient().Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	var result = make(map[string]string)
	if err = json.Unmarshal(data, &result); err != nil {
		return "", "", err
	}
	cookie := strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]
	return result["result"], cookie, nil
}

func getClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

func getClientWithRedirect() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 1 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
}
