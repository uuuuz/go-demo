package main

import (
	"encoding/json"
	"fmt"
	"github.com/mattn/go-isatty"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	//dcm "github.com/shopspring/decimal"

	"github.com/tealeg/xlsx"
)

func Test_ENV(t *testing.T) {
	fmt.Println(os.Getenv("CAM_ENV"))
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pwd)

	fmt.Println(filepath.Join(pwd, "back", "config-stage.yml"))
}

func Test_cal(t *testing.T) {
	fmt.Println("111")

	// str := ``
	// fmt.Println(str)
	out := gin.DefaultWriter
	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		// isTerm = false
		fmt.Println("yes")
	}
}

type JSON map[string]interface{}

type Decimal struct {
	value *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	exp int32
}

type ZipmexInterestDetail struct {
	RecordId     string    `db:"record_id" json:"record_id"`
	Department   string    `db:"department" json:"department"`
	BaseDate     time.Time `db:"base_date" json:"base_date"`
	Direction    string    `db:"direction" json:"direction"`
	Counterparty string    `db:"counterparty" json:"counterparty"`
	BaseCoin     string    `db:"base_coin" json:"base_coin"`
	Count        Decimal   `db:"count" json:"count"`
	Price        Decimal   `db:"price" json:"price"`
	Remark       string    `db:"remark" json:"remark"`
}

func Test_json(t *testing.T) {
	str := `{
    "department": "部门2",
    "date_time": "2021-11-05T20:03:22.502604+08:00",
    "direction": "income",
    "asset_base": "BTC",
    "counterparty_type": "customer",
    "amount": "0.01",
    "price": "60000.0",
    "remark": "this is test"
}
`

	//var cjson JSON
	//err := json.Unmarshal([]byte(str), &cjson)
	//if err != nil { //json 格式错误？
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(cjson)

	var zid ZipmexInterestDetail
	err := json.Unmarshal([]byte(str), &zid)
	if err != nil { //json 格式错误？
		fmt.Println(err.Error())
	}
	fmt.Println(zid)

	fmt.Println(time.Now().UnixNano())
}

func Test_defer(t *testing.T) {
	// test1()

	// time.Sleep(time.Second * 10)

	//fmt.Println(time.Now().AddDate(0, 0, -7).Format(time.RFC3339Nano))

	fmt.Println(assert.Contains(t, "Hello World", "World"))
	//assert.Contains(t, ["Hello", "World"], "World")
	fmt.Println(assert.Contains(t, []string{"Hello", "World"}, "Hello"))

	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().AddDate(0, 0, -7).UnixNano())
}

func test1() {
	name := "marisa"
	defer func() {
		go sleep(name)
	}()
	name = "uuz"
}

func sleep(name string) {
	time.Sleep(time.Second * 5)
	fmt.Println(name)
}

func Test_cal1(t *testing.T) {
	fmt.Println(18 * 15)
	fmt.Println(16 * 16)
}

func Test_server(t *testing.T) {
	r := gin.Default()

	r.POST("/otc-trade/manual-interest/import", func(ctx *gin.Context) {
		defer ctx.JSON(200, `{"code":"ok"}`)
		of, fh, err := ctx.Request.FormFile("file")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//mr, err := ctx.Request.MultipartReader()
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//mr.NextPart()
		//bytes.bytes
		defer of.Close()
		f, err := xlsx.OpenReaderAt(of, fh.Size)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		out, err := f.ToSlice()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for i := range out {
			fmt.Println(out[i])
		}
		return
	})

	if err := http.ListenAndServe(":8001", r); err != nil {
		fmt.Println(err.Error())
	}
}

func Test_jwt(t *testing.T) {
	//jwt.ParseECPrivateKeyFromPEM()
}

func Test_timeZone(t *testing.T) {
	fmt.Println(time.Now().Location())

	tt, _ := time.Parse("2006-01-02 15:04:05", "2021-11-14 12:34:34")
	fmt.Println(tt.Unix())
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-11-14 12:34:34", time.UTC)
	fmt.Println(t1.Unix())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-11-14 12:34:34", time.Local)
	fmt.Println(t2.Unix())
	fmt.Println(time.Now())

	t3, err := time.Parse("2006/01/02", "2021/12/23")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(t3)

	t4, err := time.Parse(time.RFC3339Nano, "2021-08-31T19:42:50.393000+08:00")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(t4)

	t5, err := time.Parse("2006-01-02 15:04:05", "2021-08-31")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(t5)
}

func Test_m(t *testing.T) {
	//fmt.Println(123456789 / 1000)
	fmt.Println(time.Now().AddDate(0, 0, -2).UnixNano())
	fmt.Println(time.Now().UnixNano())
	//fmt.Println(time.Now().Format(time.RFC3339Nano))

	fmt.Println(time.Unix(0, 1636905600000000000))
	fmt.Println(time.Unix(0, 1636992000000000000))

	fmt.Println(time.Unix(0, 1636905600000000000))

	fmt.Println(time.Now().AddDate(0, -7, -10).Format("2006/1/2"))

	fmt.Println(time.Unix(0, 1636905600000000000))
	fmt.Println(time.Unix(0, 1636905600000000000).UTC())
	//

	tz, err := strconv.Atoi(strings.Split("+4:00", ":")[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tz)

	fmt.Println(time.Unix(0, 1637110800000000000))

	//1636905600000000000
	//1636992000000000000
	//1636992000000000000
	fmt.Println(time.Unix(0, 1636992000000000000))

	//fmt.Println(time.Unix(0, 1636905600000000000))
	fmt.Println(time.Unix(0, 1636905600000000000))

}

func Test_calca(t *testing.T) {

	fmt.Println(31995296849.1 - 31451927708.9) // 31,095,980,570.1  30,568,715,734.1
	fmt.Println(540809140.4)                   // 524,367,623.8
}

type User struct {
	Name    string `db:"name"`
	Age     int    `db:"age"`
	Field1  string `db:"field1"`
	Field2  string `db:"field2"`
	Field3  string `db:"field3"`
	Field4  string `db:"field4"`
	Field5  string `db:"field5"`
	Field6  string `db:"field6"`
	Field7  string `db:"field7"`
	Field8  string `db:"field8"`
	Field9  string `db:"field9"`
	Field10 string `db:"field10"`
	Field11 string `db:"field11"`
	Field12 string `db:"field12"`
	Field13 string `db:"field13"`
	Field14 string `db:"field14"`
	Field15 string `db:"field15"`
	Field16 string `db:"field16"`
	Field17 string `db:"field17"`
	Field18 string `db:"field18"`
}

func Test_reflect(t *testing.T) {
	start := time.Now().UnixNano()
	defer func() {
		fmt.Println(time.Now().UnixNano() - start)
	}()
	var dest interface{} = &User{}
	check := reflect.TypeOf(dest).Elem()
	fields := make([]string, 0)
	for i := 0; i < check.NumField(); i++ {
		db := check.Field(i).Tag.Get("db")
		if db != "" {
			fields = append(fields, db)
		}
	}

}

func Test_nil(t *testing.T) {
	direction := ""
	str := strings.Split(direction, ",")
	fmt.Println(str)
	fmt.Println(str == nil)
	fmt.Println(len(str) == 0)

	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().AddDate(0, 0, -1).UnixNano())
}

var sw sync.WaitGroup

func Test_http(t *testing.T) {
	for i := 0; i < 100; i++ {
		sw.Add(1)
		go goJava(i)
	}
	sw.Wait()
}

func goJava(n int) {
	defer sw.Done()
	for i := 0; i < 10; i++ {
		//http.Get("http://localhost:8000/test1")
		//http.Get("http://localhost:30001/call/get")
		http.Get("http://47.103.121.55:30001/call/get")
		//http.Get("http://localhost:10001/call/other")
		fmt.Println(n, i)
	}
}

func Test_time11(t *testing.T) {
	fmt.Println(time.Unix(0, 1577808000000000000))
}

type MyString string

func Test_str(t *testing.T) {
	var str MyString = "all"
	fmt.Println(str == "all")
	param := make(map[string]string)
	res := optMap(param)
	fmt.Println(param)
	fmt.Println(res)
}

func optMap(param map[string]string) map[string]string {
	param["test"] = "test"
	return param
}

func Test_char(t *testing.T) {
	//fmt.Println(strings.)
	builder := strings.Builder{}
	if err := builder.WriteByte(90); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(builder.String())
	fmt.Println(string([]byte{90}))
}

func Test_getSheetABCColByNum(t *testing.T) {
	fmt.Println(getSheetABCColByNum(2))
	fmt.Println(getSheetABCColByNum(25))
	fmt.Println(getSheetABCColByNum(26))
	fmt.Println(getSheetABCColByNum(27))
	fmt.Println(getSheetABCColByNum(106))
}

func getSheetABCColByNum(num int) string {
	// 65 - 90
	n, r := num/26, num%26
	if n == 0 {
		return string([]byte{byte(r + 65)})
	}
	return getSheetABCColByNum(n-1) + string([]byte{byte(r + 65)})
}

func Test_timeout(t *testing.T) {
	cli := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := cli.Get("http://localhost:1111/test")
	assert.Nil(t, err)
	t.Log(resp)
}

func TestMap(t *testing.T) {
	m := make(map[string]bool)
	m[""] = true
	m["a"] = true
	m["b"] = true

	assert.True(t, m[""])
	assert.True(t, m["a"])
	assert.True(t, !m["c"])
}

func TestFloat(t *testing.T) {
	expiryOfYear := time.Until(time.Now().AddDate(0, 0, 2)).Hours() / 24 / 365
	t.Log(expiryOfYear)
}
