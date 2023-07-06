package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func Test_img(t *testing.T) {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{
    "type":"doughnut",
    "series":[
        {
            "name":"test it",
            "categories":"Sheet1!$B$1:$D$1",
            "values":"Sheet1!$B$2:$D$2"
        }
    ],
    "format":{
        "x_scale":1,
        "y_scale":1,
        "x_offset":15,
        "y_offset":10,
        "print_obj":true,
        "lock_aspect_ratio":false,
        "locked":false
    },
    "legend":{
        "position":"right",
        "show_legend_key":false
    },
    "title":{
        "name":"Doughnut Chart"
    },
    "plotarea":{
        "show_bubble_size":false,
        "show_cat_name":false,
        "show_leader_lines":false,
        "show_percent":true,
        "show_series_name":false,
        "show_val":false
    },
    "show_blanks_as":"zero"
}`); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetSheetViewOptions("Sheet1", -1, excelize.ShowGridLines(false)); err != nil {
		fmt.Println(err)
		return
	}
	//f.SetSheetFormatPr("Sheet1")
	//f.SetSheetPrOptions()
	// Save spreadsheet by the given path.
	if err := f.SaveAs("./Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func Test_it(t *testing.T) {
	file, sheetName := excelize.NewFile(), "Sheet1"
	categories := map[string]string{"A1": "Apple", "B1": "Orange", "C1": "Pear"}
	values := map[string]int{"A2": 2, "B2": 3, "C2": 3}

	var err error
	for k, v := range categories {
		err = file.SetCellValue(sheetName, k, v)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	for k, v := range values {
		err = file.SetCellValue(sheetName, k, v)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	err = file.SetSheetViewOptions(sheetName, -1, excelize.ShowGridLines(false))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = file.AddChart("Sheet1", "E1", `{
		   "type": "doughnut",
		   "series": [
		   {
		       "name": "series 1",
		       "categories": "Sheet1!$A$1:$C$1",
		       "values": "Sheet1!$A$2:$C$2"
		   }],
		   "format":
		   {
		       "x_scale": 1.0,
		       "y_scale": 1.0,
		       "x_offset": 15,
		       "y_offset": 10,
		       "print_obj": true,
		       "lock_aspect_ratio": false,
		       "locked": false
		   },
		   "legend":
		   {
		       "position": "right",
		       "show_legend_key": false
		   },
		   "title":
		   {
		       "name": ""
		   },
		   "plotarea":
		   {
		       "show_bubble_size": false,
		       "show_cat_name": false,
		       "show_leader_lines": false,
		       "show_percent": true,
		       "show_series_name": false,
		       "show_val": false
		   },
		   "show_blanks_as": "zero"
		}`); err != nil {

		fmt.Println(err.Error())
		return
	}

	err = file.SaveAs("test.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
