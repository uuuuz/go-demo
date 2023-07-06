package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"testing"
)

func Test_calcPivotTable(t *testing.T) {
	// 生成原数据
	data := genOriData()
	// excel生成透视表
	if err := genPivotTable(data); err != nil {
		fmt.Println(err)
		return
	}
	// 计算生成透视表
	title := make([]string, len(data[0]))
	for i := range title {
		title[i] = fmt.Sprintf("%v", data[0][i])
	}
	body := make([][]string, len(data)-1)
	for i, row := range data[1:] {
		line := make([]string, len(row))
		for j := range line {
			line[j] = fmt.Sprintf("%v", row[j])
		}
		body[i] = line
	}
	res, err := CalcPivotTable(title, body, PivotTableOption{
		Filter: []PivotTableFilterOption{{ColumnName: "Region"}},
		Rows:   []string{"Year", "Month"},
		//Columns:    []string{"Type"},
		Values:     []string{"Sales1"}, // Sales1
		ValueInRow: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = genExcel(res); err != nil {
		fmt.Println(err)
		return
	}
}

func genOriData() [][]interface{} {
	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	year := []int{2017, 2018, 2019}
	types := []string{"Meat", "Dairy", "Beverages", "Produce"}
	region := []string{"East", "West", "North", "South"}

	title := []interface{}{"Month", "Year", "Type", "Region", "Sales1", "Sales2"}
	var data [][]interface{}
	data = append(data, title)
	for row := 2; row < 32; row++ {
		data = append(data, []interface{}{month[rand.Intn(12)], year[rand.Intn(3)], types[rand.Intn(4)], region[rand.Intn(4)], rand.Intn(5000), rand.Intn(5000)})
	}
	return data
}

func genPivotTable(data [][]interface{}) error {
	f := excelize.NewFile()
	defer f.Close()
	for i := range data {
		for j, v := range data[i] {
			cell, err := excelize.CoordinatesToCellName(j+1, i+1)
			if err != nil {
				return err
			}
			if err = f.SetCellValue("Sheet1", cell, v); err != nil {
				return err
			}
		}
	}
	if err := f.AddPivotTable(&excelize.PivotTableOption{
		DataRange:       "Sheet1!$A$1:$F$31",
		PivotTableRange: "Sheet1!$H$2:$Z$35",
		Rows: []excelize.PivotTableField{
			{Data: "Type"}},
		Columns: []excelize.PivotTableField{
			{Data: "Year"}, {Data: "Month"}},
		Data: []excelize.PivotTableField{
			{Data: "Sales1", Name: "Summarize1", Subtotal: "Sum"}, {Data: "Sales2", Name: "Summarize2", Subtotal: "Sum"}}, // {Data: "Sales1", Name: "Summarize1", Subtotal: "Sum"},
		Filter: []excelize.PivotTableField{
			{Data: "Region"}},
		RowGrandTotals: true,
		ColGrandTotals: true,
		ShowDrill:      true,
		ShowRowHeaders: true,
		ShowColHeaders: true,
		ShowLastColumn: true,
	}); err != nil {
		return err
	}
	if err := f.SaveAs("/Users/wxm/Desktop/Book1.xlsx"); err != nil {
		return err
	}
	return nil
}

func genExcel(data [][]interface{}) error {
	f := excelize.NewFile()
	defer f.Close()
	for i := range data {
		for j, v := range data[i] {
			cell, err := excelize.CoordinatesToCellName(j+1, i+1)
			if err != nil {
				return err
			}
			if err = f.SetCellValue("Sheet1", cell, v); err != nil {
				return err
			}
		}
	}
	if err := f.SaveAs("/Users/wxm/Desktop/Book2.xlsx"); err != nil {
		return err
	}
	return nil
}
