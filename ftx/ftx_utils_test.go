package ftx

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const (
	APIKEY    = "vtqDy_Z4i_6ruYDTErV3kHI3Wek_c5wym8D41LWM"
	APISECRET = "AL3Tp_vA1On6f188t0MsDoYr9RHqpmfkLN7H3w2q"
)

func Test_run(t *testing.T) {
	// 组装头部
	method, path := "GET", "/wallet/balances"
	req, err := http.NewRequest(method, "https://ftx.com/api"+path, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// ts
	ts := fmt.Sprintf("%d", time.Now().Unix()*1000)
	// sign
	//h := hmac.New(sha256.New, []byte(APISECRET))
	//t.Log(ts + method + path)
	//h.Write([]byte(ts + method + path))
	//sign := hex.EncodeToString(h.Sum(nil))
	
	head := http.Header{}
	head.Add("FTX-KEY", APIKEY)
	head.Add("FTX-SIGN", "00b2b819b595813de0ba99df22004193a9b826efd2288989ce98b88277bb0e5b")
	//head.Add("FTX-SIGN", sign)
	head.Add("FTX-TS", ts)
	req.Header = head
	t.Log(head)
	// client
	proxyURL, err := url.Parse("http://ubuntu.urwork.qbtrade.org:1080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			//Proxy: http.ProxyFromEnvironment,  // 使用环境变量代理
			Proxy: http.ProxyURL(proxyURL), // 使用传入的URL代理
		},
	}
	// sent
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Log(resp)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Log(string(data))
}
