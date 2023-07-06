package main

import (
	"encoding/base64"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"testing"
	"time"
)

func Test_ReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("/Users/wxm/Desktop/工作簿1.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := base64.StdEncoding.EncodeToString(data)
	fmt.Println(res)
}

func Test_slice(t *testing.T) {
	s := []string{"1"}
	//s1 := s[1:]
	fmt.Println(len(s[0:]) == 0)
}

func Test_time(t *testing.T) {
	//fmt.Println(time.Now().AddDate(0, 0, -10).UnixNano())
	//fmt.Println(time.Now().UnixNano())

	ti, err := time.ParseInLocation("2006-01-02", "2021-11-10", time.Local)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ti)

	time.ParseInLocation("2006-01-02", "2021-11-10", time.Local)

	start, _ := time.ParseInLocation("2006-01-02", "2021-11-01", time.Local)
	end, _ := time.ParseInLocation("2006-01-02", "2021-11-11", time.Local)
	end = end.Add(-1)
	fmt.Println("start: ", start)
	fmt.Println("end: ", end)

	start, _ = time.ParseInLocation("2006-01", "2021-10", time.Local)
	end, _ = time.ParseInLocation("2006-01", "2021-12", time.Local)
	end = end.Add(-1)
	fmt.Println("start: ", start)
	fmt.Println("end: ", end)

	fmt.Println(time.Now().UnixNano()) // 1636629165856610000
	fmt.Println(time.Unix(1636629165, 856610000))

	tn := time.Now()
	fmt.Println(tn)
	fmt.Println(time.Unix(tn.Unix(), tn.UnixNano()%1000000000))

	fmt.Println(time.Now().AddDate(0, -4, 0))
	fmt.Println(time.Now().AddDate(0, -4, 0).UnixNano())
}

func Test_readAndCheckExcel(t *testing.T) {
	// 创建xlsx文件，并设置列格式
	f := xlsx.NewFile()
	s, _ := f.AddSheet("01")
	// 新增行
	r := s.AddRow()
	// 新增单元格
	c1 := r.AddCell()
	// 设置单元格值
	c1.SetString("2021-11-12 12:12:35")
	// 设置单元格样式
	style := xlsx.NewStyle()
	// 设置字体
	style.Font = *xlsx.NewFont(20, "黑体")
	// 设置边框样式
	style.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	// 填充
	style.Fill = *xlsx.NewFill(xlsx.Solid_Cell_Fill, xlsx.RGB_Light_Red, xlsx.RGB_White)
	// 对齐
	style.Alignment = xlsx.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	}
	xlsx.DefaultAlignment()

	c1.SetStyle(style)
	if err := f.Save("./test.xlsx"); err != nil {
		fmt.Println(err.Error())
	}
}

func Test_fileServer(t *testing.T) {
	//http.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
	//	// todo download file
	//})
	//
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	fmt.Println(err.Error())
	//}

	ti := time.Now().AddDate(0, 1, 0)
	y, m, d := ti.Date()
	startTime := time.Date(y, m, 1, 0, 0, 0, 0, ti.Location())
	startTime1 := time.Date(y, m, d, 0, 0, 0, 0, ti.Location())

	fmt.Println(ti)
	fmt.Println(startTime)
	fmt.Println(startTime1)

	// ioutil.NopCloser()
}

func Test_createFile(t *testing.T) {
	f := xlsx.NewFile()
	s, _ := f.AddSheet("01")
	// 新增行
	sumRow := s.AddRow()
	sumRow.AddCell().SetFormula(fmt.Sprintf("SUM(A2:A4)"))
	r := s.AddRow()
	r.AddCell().SetFloat(1.22)
	r = s.AddRow()
	r.AddCell().SetFloat(1.22)
	r = s.AddRow()
	r.AddCell().SetFloat(1.22)

	if err := f.Save("./test.xlsx"); err != nil {
		fmt.Println(err.Error())
	}
}

func Test_time1(t *testing.T) {
	ti, err := time.ParseInLocation("2006-01-02", "2022-10-21", time.Local)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ti2 := time.Date(ti.Year(), ti.Month()+1, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(ti2.AddDate(0, 0, -1).Format("2006-01-02"))
}

func Test_time2(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().AddDate(0, -2, 0).UnixNano())
}
