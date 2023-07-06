package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"math/rand"
)

func main() {
	if err := genPivotTableData(); err != nil {
		fmt.Println(err)
	}
	//if err := readPivotTable(); err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println("over")
}

func genPivotTableData() error {
	f := excelize.NewFile()
	// Create some data in a sheet
	month := []string{"Jan", "Feb", "Mar", "Apr", "May",
		"Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	year := []int{2017, 2018, 2019}
	types := []string{"Meat", "Dairy", "Beverages", "Produce"}
	region := []string{"East", "West", "North", "South"}
	if err := f.SetSheetRow("Sheet1", "A1", &[]string{"Month", "Year", "Type", "Sales", "Region"}); err != nil {
		return err
	}
	for row := 2; row < 32; row++ {
		if err := f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), month[rand.Intn(12)]); err != nil {
			return err
		}
		if err := f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), year[rand.Intn(3)]); err != nil {
			return err
		}
		if err := f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), types[rand.Intn(4)]); err != nil {
			return err
		}
		if err := f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), rand.Intn(5000)); err != nil {
			return err
		}
		if err := f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), region[rand.Intn(4)]); err != nil {
			return err
		}
	}
	f.NewSheet("Sheet2")
	if err := f.AddPivotTable(&excelize.PivotTableOption{
		DataRange:       "Sheet1!$A$1:$E$31",
		PivotTableRange: "Sheet1!$G$2:$M$35",
		//PivotTableRange: "Sheet2!$A$1:$G$33",
		Rows: []excelize.PivotTableField{
			{Data: "Month", DefaultSubtotal: true}, {Data: "Year"}},
		Columns: []excelize.PivotTableField{
			{Data: "Type", DefaultSubtotal: true}},
		Data: []excelize.PivotTableField{
			{Data: "Sales", Name: "Summarize", Subtotal: "Sum"}},
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
		fmt.Println(err)
	}
	if err := f.Close(); err != nil {
		return err
	}
	//data, err := f.GetRows(f.GetSheetName(0))
	//if err != nil {
	//	return err
	//}
	//formatData, err := json.Marshal(data)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("data1:", string(formatData))
	//fmt.Println("------------")
	//data, err = f.GetRows(f.GetSheetName(1))
	//if err != nil {
	//	return err
	//}
	//formatData, err = json.Marshal(data)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("data2:", string(formatData))
	//fmt.Println("------------")
	return nil
}

func readPivotTable() error {
	data, err := ioutil.ReadFile("/Users/wxm/Desktop/Book1.xlsx")
	if err != nil {
		return err
	}
	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer f.Close()

	sheetName := f.GetSheetName(1)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	formatData, err := json.Marshal(rows)
	if err != nil {
		return err
	}
	fmt.Println("data:", string(formatData))
	fmt.Println("------------")

	v, err := f.GetCellValue(sheetName, "B5")
	if err != nil {
		return err
	}
	fmt.Println(v)
	return nil
}
