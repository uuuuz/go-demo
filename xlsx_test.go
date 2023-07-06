package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"testing"
)

func Test_File(t *testing.T) {
	f := xlsx.NewFile()
	s, _ := f.AddSheet("01")

	// 填充数据
	row := s.AddRow()
	row.AddCell().Value = "01"
	row.AddCell().Value = "02"
	row.AddCell().Value = "03"
	row = s.AddRow()
	row.AddCell().Value = "01"
	row.AddCell().Value = "02"
	row.AddCell().Value = "03"

	// 设置表格
	style := xlsx.NewStyle()
	// font := *xlsx.NewFont(12, "Verdana")
	//font.Bold = true
	//font.Italic = true
	//font.Underline = true
	//style.Font = font
	//fill := *xlsx.NewFill("solid", "00FF0000", "FF000000")
	//style.Fill = fill
	//style.ApplyFill = true
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Border = border
	style.ApplyBorder = true

	row = s.AddRow()
	row.AddCell()
	row.AddCell()
	row = s.AddRow()
	row.AddCell()
	row.AddCell()

	cells := s.Row(0).Cells
	style = xlsx.NewStyle()
	style.Border = *xlsx.NewBorder("thin", "", "thin", "")
	cells[0].SetStyle(style)

	style = xlsx.NewStyle()
	style.Border = *xlsx.NewBorder("", "", "thin", "")
	cells[1].SetStyle(style)

	style = xlsx.NewStyle()
	style.Border = *xlsx.NewBorder("", "thin", "thin", "")
	cells[2].SetStyle(style)

	cells = s.Row(1).Cells
	style = xlsx.NewStyle()
	style.Border = *xlsx.NewBorder("thin", "", "", "thin")
	cells[0].SetStyle(style)

	style = xlsx.NewStyle()
	style.Border = *xlsx.NewBorder("", "", "", "thin")
	style = xlsx.NewStyle()
	style.Border = *xlsx.NewBorder("", "thin", "", "thin")
	cells[2].SetStyle(style)
	// 去除网格
	// todo
	// 填充
	// todo
	if err := f.Save("./test.xlsx"); err != nil {
		fmt.Println(err.Error())
	}
}

func Test_grid(t *testing.T) {
	f := xlsx.NewFile()
	s, _ := f.AddSheet("01")

	// 填充数据
	row := s.AddRow()
	row.AddCell().Value = "01"
	row.AddCell().Value = "02"
	row.AddCell().Value = "03"
	row = s.AddRow()
	row.AddCell().Value = "01"
	row.AddCell().Value = "02"
	row.AddCell().Value = "03"
	// 去除网格
	ap := s.SheetViews[0].Pane.ActivePane
	t.Log(ap)
	if err := f.Save("./test.xlsx"); err != nil {
		fmt.Println(err.Error())
	}
}
