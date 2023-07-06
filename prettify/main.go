package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pterm/pterm"
	"os"
)

func main() {
	// github.com/jedib0t/go-pretty/v6/table
	createDB1()
	//createDB2()
}

func createDB1() {
	columns, data := getData()
	hr := make(table.Row, len(columns))
	for index, v := range columns {
		hr[index] = v
	}
	t := table.NewWriter()
	rcAutoMerge := table.RowConfig{AutoMerge: true}
	// 设置列名
	t.AppendHeader(hr, rcAutoMerge)
	//t.AppendHeader(hr)
	for _, items := range data {
		dr := make(table.Row, len(items))
		for i, v := range items {
			dr[i] = wrapHard(v, 40)
			//dr[i] = v
		}
		t.AppendRow(dr, rcAutoMerge)
		//t.AppendRow(dr)
		t.AppendSeparator()
	}
	//colConfigs := make([]table.ColumnConfig, 0)
	//colConfigs = append(colConfigs, table.ColumnConfig{Number: 1, WidthMax: 40, WidthMaxEnforcer: text.WrapHard})
	//colConfigs = append(colConfigs, table.ColumnConfig{Number: 2, WidthMax: 40, WidthMaxEnforcer: text.WrapHard})
	//colConfigs = append(colConfigs, table.ColumnConfig{Number: 3, WidthMax: 40, WidthMaxEnforcer: text.WrapHard})
	//colConfigs = append(colConfigs, table.ColumnConfig{Number: 4, WidthMax: 40, WidthMaxEnforcer: text.WrapHard})
	//t.SetColumnConfigs(colConfigs)
	//t.SetAllowedRowLength(100)
	f, err := os.Create("./test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	t.SetOutputMirror(f)
	//t.SetOutputMirror(os.Stdout)
	//t.SetAutoIndex(true)
	//t.Style().Options.SeparateRows = true
	//t.Render()
	t.Render()
}

func wrapHard(v string, max int) string {
	data := []rune(v)
	rs := make([]rune, 0, len(data)+len(data)/max+1)
	for i, e := range data {
		rs = append(rs, e)
		if i%max == max-1 {
			rs = append(rs, '\n')
		}
	}
	return string(rs)
}

func createDB2() {
	columns, data := getData2()
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		columns,
		data[0],
		data[1],
		//data[2],
		//data[3],
		//data[4],
	}).Render()
}

func getData() ([]string, [][]string) {
	return []string{
			"column_1", "column_2", "column_3", "column_4", "column_5", "column_6", "column_7", "column_8", "column_9", "column_10", "column_11", "column_12", "column_13", "column_14", "column_15", "column_16", "column_17", "column_18", "column_19", "column_20",
		}, [][]string{
			{
				"value_1", "value_2", "value_3", "value_4", "value_5", "value_6", "value_7", "value_8", "value_9", "value_10", "value_11", "value_12", "value_13", "value_14", "value_15", "value_16", "value_17", "value_18", "value_19", "value_20",
			},
			//			{
			//				`quote目录下的camquote_service
			//包提供了一些查询历史行情的方法，
			//目前这些方法支持查询cmc和index2
			//的历史行情(cmc支持2021年1月1日之
			//后的每小时k线和2020年1月1日之后
			//的每日k线，index2 dev环境有
			//2022年1月1日之后的每分钟k线，
			//prod环境是2021年1月1日之后的
			//每分钟k线)。 `, "value_2", "value_3", "value_4", "value_5", "value_6", "value_7", "value_8", "value_9", "value_10", "value_11", "value_12", "value_13", "value_14", "value_15", "value_16", "value_17", "value_18", "value_19", "value_20",
			//			},
			{
				"value_1", "value_2", "value_3", "value_4", "value_5", "value_6", "value_7", "value_8", "value_9", "value_10", "value_11", "value_12", "value_13", "value_14", "value_15", "value_16", "value_17", "value_18", "value_19", "value_20",
			},
			{
				"value_1", "现代的电子计算器能进行数学运算的手持电子机器，拥有集成电路芯片，但结构比电脑/usr/local/Cellar/go@1.18/1.18.6/libexec/bin/go build -o /private/var/folders/ll/0pm6bc3j2bb6ln82gplsv1v80000gn/T/GoLand/___1go_b简单得多，可以说是第一代的电子计算机（电脑），且功能也较弱，但较为方便/usr/local/Cellar/go@1.18/1.18.6/libexec/bin/go build -o /private/var/folders/ll/0pm6bc3j2bb6ln82gplsv1v80000gn/T/GoLand/___1go_b与廉价，可广泛运用于商业交易中，是必备的办公用品之一。除显示计算结果外，还常有溢出/usr/local/Cellar/go@1.18/1.18.6/libexec/bin/go build -o /private/var/folders/ll/0pm6bc3j2bb6ln82gplsv1v80000gn/T/GoLand/___1go_b指示、错误指示等。计算器电源采用交流转换器或电池，电池可用交流转换器或太阳能转换器再充电。为节省电能，计算器都采用CMOS工艺制作的大规模集成电路。", "value_3", "value_4", "value_5", "value_6", "value_7", "value_8", "value_9", "value_10", "value_11", "value_12", "value_13", "value_14", "value_15", "value_16", "value_17", "value_18", "value_19", "value_20",
			},
			{
				"value_1", "value_2", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "value_4", "value_5", "value_6", "value_7", "value_8", "value_9", "value_10", "value_11", "value_12", "value_13", "value_14", "value_15", "value_16", "value_17", "value_18", "value_19", "value_20",
			},
			//{
			//	"value_1", "value_2", "value_3", "value_4", "value_5", "value_6", "value_7", "value_8", "value_9", "value_10", "value_11", "value_12", "value_13", "value_14", "value_15", "value_16", "value_17", "value_18", "value_19", "value_20",
			//},
			//{
			//	"value_1", "value_2", "value_3", "value_4", "value_5", "value_6", "value_7", "value_8", "value_9", "value_10", "value_11", "value_12", "value_13", "value_14", "value_15", "value_16", "value_17", "value_18", "value_19", "value_20",
			//},
		}
}

func getData2() ([]string, [][]string) {
	return []string{
			"column_1", "column_2", "column_3", "column_4",
		}, [][]string{
			//			{
			//				"value_1", "value_2", "value_3", `quote目录下的camquote_service
			//包提供了一些查询历史行情的方法，
			//目前这些方法支持查询cmc和index2
			//的历史行情(cmc支持2021年1月1日之
			//后的每小时k线和2020年1月1日之后
			//的每日k线，index2 dev环境有
			//2022年1月1日之后的每分钟k线，
			//prod环境是2021年1月1日之后的
			//每分钟k线)。 `,
			//			},
			{
				"value_1", "现代的电子计算器能进行数学运算的手持电子机器，拥有集成电路芯片，但结构比电脑/usr/local/Cellar/go@1.18/1.18.6/libexec/bin/go build -o /private/var/folders/ll/0pm6bc3j2bb6ln82gplsv1v80000gn/T/GoLand/___1go_b简单得多，可以说是第一代的电子计算机（电脑），且功能也较弱，但较为方便/usr/local/Cellar/go@1.18/1.18.6/libexec/bin/go build -o /private/var/folders/ll/0pm6bc3j2bb6ln82gplsv1v80000gn/T/GoLand/___1go_b与廉价，可广泛运用于商业交易中，是必备的办公用品之一。除显示计算结果外，还常有溢出/usr/local/Cellar/go@1.18/1.18.6/libexec/bin/go build -o /private/var/folders/ll/0pm6bc3j2bb6ln82gplsv1v80000gn/T/GoLand/___1go_b指示、错误指示等。计算器电源采用交流转换器或电池，电池可用交流转换器或太阳能转换器再充电。为节省电能，计算器都采用CMOS工艺制作的大规模集成电路。", "value_3", "value_4",
			},
			{
				"value_1", "value_2", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR",
			},
			//{
			//	"quote目录下的camquote_service包提供了一些查询历史行情的方法，目前这些方法支持查询cmc和index2的历史行情(cmc支持2021年1月1日之后的每小时k线和2020年1月1日之后的每日k线，index2 dev环境有2022年1月1日之后的每分钟k线，prod环境是2021年1月1日之后的每分钟k线)。 ", "value_2", "value_3", "value_4",
			//},
			//{
			//	"value_1", "value_2", "quote目录下的camquote_service包提供了一些查询历史行情的方法，目前这些方法支持查询cmc和index2的历史行情(cmc支持2021年1月1日之后的每小时k线和2020年1月1日之后的每日k线，index2 dev环境有2022年1月1日之后的每分钟k线，prod环境是2021年1月1日之后的每分钟k线)。 ", "value_4",
			//},
		}
}
