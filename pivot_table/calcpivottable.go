package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type PivotTableFilterOption struct {
	ColumnName string
	Items      []string // 为空则全部
}

type PivotTableOption struct {
	Filter     []PivotTableFilterOption
	Rows       []string
	Columns    []string
	Values     []string
	ValueInRow bool // false 在列标签上，仅 len(Values) > 1 时生效
}

type tagInfo struct {
	values []string
}

type condition struct {
	index int // 条件索引
	value string
}

type rowInfo struct {
	items []interface{}
}

func CalcPivotTable(title []string, data [][]string, option PivotTableOption) ([][]interface{}, error) {
	// 得到列名与索引位置的映射关系
	columnIndexMap := make(map[string]int)
	for i, val := range title {
		columnIndexMap[val] = i
	}
	// 根据 filter 过滤数据
	data = filterData(columnIndexMap, data, option.Filter)
	// 生成列标签
	columnTags := genTags(columnIndexMap, data, option.Columns)
	// 生成行标签
	rowTags := genTags(columnIndexMap, data, option.Rows)
	// 分组并汇总求和
	res := genData(columnIndexMap, columnTags, rowTags, data, option.Columns, option.Rows, option.Values, option.ValueInRow)
	return res, nil
}

func genData(columnIndexMap map[string]int, columnTags, rowTags []tagInfo, data [][]string, condColumns, condRows, values []string, inRow bool) [][]interface{} {
	if len(columnTags) == 0 && len(rowTags) == 0 {
		return make([][]interface{}, 0)
	}
	// gen frame
	frame := genFrame(columnIndexMap, columnTags, rowTags, data, condColumns, condRows, values, inRow)
	return frame
}

// 画一个矩阵，然后遍历矩阵，进行赋值
func genFrame(columnIndexMap map[string]int, columnTags, rowTags []tagInfo, data [][]string, condColumns, condRows, values []string, inRow bool) [][]interface{} {
	totalRow, totalCol := 1, 1
	if len(rowTags) == 0 {
		totalRow = 0
	}
	if len(columnTags) == 0 {
		totalCol = 0
	}
	blankCellRow, blankCellCol := len(columnTags), len(rowTags)
	if len(values) > 1 {
		if inRow {
			blankCellCol += 1
			totalRow = len(values)
			if len(rowTags) == 0 {
				totalRow = 0
			}
			rowTags = append(rowTags, tagInfo{values: values})
		} else {
			blankCellRow += 1
			totalCol = len(values)
			if len(columnTags) == 0 {
				totalCol = 0
			}
			columnTags = append(columnTags, tagInfo{values: values})
		}
	}
	// 生成矩阵 - 此时的 table 的数值都是 ""。下面的统计中不需要设置默认值
	colGroupLen, rowGroupLen, table := genMatrix(columnTags, rowTags, totalRow, totalCol)
	// 记录需要删除的列
	needSkipColMap := make(map[int]bool)
	for col := range table[0].items {
		needSkipColMap[col] = true
		if col < blankCellCol || col >= len(table[0].items)-totalCol {
			needSkipColMap[col] = false
		}
	}
	// 记录需要删除的行
	needSkipRowMap := make(map[int]bool)
	for row := range table {
		needSkipRowMap[row] = true
		if row < blankCellRow || row >= len(table)-totalRow {
			needSkipRowMap[row] = false
		}
	}
	// 填充矩阵
	for row := range table {
		// fill data
		for col := range table[row].items {
			if row < blankCellRow && col < blankCellCol {
				continue
			}
			// 标题
			if row < blankCellRow && col >= blankCellCol {
				// 统计列处理
				if col == len(table[row].items)-totalCol && row == 0 && len(columnTags) > 0 {
					table[row].items[col] = "Total"
					continue
				}
				// value 行特殊处理
				if len(values) > 1 && !inRow && row == blankCellRow-1 {
					index := (col - blankCellCol) % len(values)
					table[row].items[col] = values[index]
					continue
				}
				if col >= len(table[row].items)-totalCol {
					continue
				}
				// 其他行与列
				titleCNum := colGroupLen
				for index := range columnTags {
					titleCNum = titleCNum / len(columnTags[index].values)
					if index >= row {
						break
					}
				}
				if (col-blankCellCol)%colGroupLen%titleCNum == 0 {
					index := (col - blankCellCol) % colGroupLen / titleCNum % len(columnTags[row].values)
					table[row].items[col] = columnTags[row].values[index]
				}
				continue
			}
			// 标题
			if col < blankCellCol && row >= blankCellRow {
				if row == len(table)-totalRow && col == 0 && len(rowTags) > 0 { // 统计行处理
					table[row].items[col] = "Total"
					continue
				}
				// value 行特殊处理
				if len(values) > 1 && inRow && col == blankCellCol-1 {
					index := (row - blankCellRow) % len(values)
					table[row].items[col] = values[index]
					continue
				}
				if row >= len(table)-totalRow {
					continue
				}
				// 其他行与列
				titleRNum := rowGroupLen
				for index := range rowTags {
					titleRNum = titleRNum / len(rowTags[index].values)
					if index >= col {
						break
					}
				}
				if (row-blankCellRow)%rowGroupLen%titleRNum == 0 {
					index := (row - blankCellRow) % rowGroupLen / titleRNum % len(rowTags[col].values)
					table[row].items[col] = rowTags[col].values[index]
				}
				continue
			}
			// 数据填充 col >= blankCellCol && row >= blankCellRow
			// 列的总计列处理
			if col >= len(table[row].items)-totalCol {
				totalVal := 0.0
				for i := col - (len(table[row].items) - totalCol); i < len(table[0].items)-blankCellCol-totalCol; i += totalCol {
					totalVal += dealValue(table[row].items[blankCellCol+i])
				}
				table[row].items[col] = totalVal
				continue
			}
			// 行的总计列
			if row >= len(table)-totalRow {
				totalVal := 0.0
				for i := row - (len(table) - totalRow); i < len(table)-blankCellRow-totalRow; i += totalRow {
					totalVal += dealValue(table[blankCellRow+i].items[col])
				}
				table[row].items[col] = totalVal
				continue
			}
			// 其他行列
			cond, cNum, rNum := make([]condition, 0, len(condColumns)+len(condRows)), colGroupLen, rowGroupLen
			for i, name := range condColumns {
				cNum = cNum / len(columnTags[i].values)
				index := ((col - blankCellCol) / cNum) % len(columnTags[i].values)
				cond = append(cond, condition{index: columnIndexMap[name], value: columnTags[i].values[index]})
			}
			for i, name := range condRows {
				rNum = rNum / len(rowTags[i].values)
				index := ((row - blankCellRow) / rNum) % len(rowTags[i].values)
				cond = append(cond, condition{index: columnIndexMap[name], value: rowTags[i].values[index]})
			}
			index := (col - blankCellCol) % len(values)
			if len(values) > 1 && inRow {
				index = (row - blankCellRow) % len(values)
			}
			val := sumByCondition(data, cond, columnIndexMap[values[index]])
			table[row].items[col] = val
			if val != 0 {
				needSkipColMap[col], needSkipRowMap[row] = false, false
			}
		}
	}
	// 去掉所有空白行（没有数值的行） // 去掉value行
	var res = make([][]interface{}, 0, len(table))
	for i, row := range table {
		// 删除行
		if needSkipRowMap[i] {
			if i < len(table)-1 {
				for c := 0; c < blankCellCol; c++ {
					if row.items[c] != nil && len(row.items[c].(string)) > 0 {
						table[i+1].items[c] = row.items[c]
					}
				}
			}
			continue
		}
		// 删除列
		line := make([]interface{}, 0, len(row.items))
		for col, v := range row.items {
			if needSkipColMap[col] {
				if i < blankCellRow && v != nil && len(v.(string)) > 0 { // 把不为空的值往后移动
					row.items[col+1] = v
				}
				continue
			}
			line = append(line, v)
		}
		res = append(res, line)
	}
	return res
}

func genMatrix(columnTags, rowTags []tagInfo, totalRow, totalCol int) (int, int, []rowInfo) {
	var colGroupLen, rowGroupLen int
	// 列数
	colLen := 1
	for _, items := range columnTags {
		colLen = colLen * len(items.values)
	}
	colGroupLen = colLen
	colLen = colLen + len(rowTags) + totalCol
	// 行数
	rowLen := 1
	for _, items := range rowTags {
		rowLen = rowLen * len(items.values)
	}
	rowGroupLen = rowLen
	rowLen = rowLen + len(columnTags) + totalRow
	// 表
	table := make([]rowInfo, rowLen)
	for i := range table {
		table[i] = rowInfo{items: make([]interface{}, colLen)}
	}
	return colGroupLen, rowGroupLen, table
}

func dealValue(val interface{}) float64 {
	var v float64
	switch val.(type) {
	case float32:
		v = float64(val.(float32))
	case float64:
		v = val.(float64)
	default:
		v, _ = strconv.ParseFloat(fmt.Sprintf("%v", val), 64)
	}
	return v
}

func sumByCondition(data [][]string, conditions []condition, sumValueIndex int) float64 {
	total := 0.0
	for _, items := range data {
		pass := true
		for _, cond := range conditions {
			if items[cond.index] != cond.value {
				pass = false
				break
			}
		}
		if pass {
			// 忽略 err todo
			v, err := strconv.ParseFloat(items[sumValueIndex], 64)
			if err != nil {
				fmt.Println(err.Error())
			}
			total += v
		}
	}
	// 精度考虑
	return total
}

func genTags(columnIndexMap map[string]int, data [][]string, group []string) []tagInfo {
	tags := make([]tagInfo, len(group))
	for i := range group {
		uniqueMap := make(map[string]struct{})
		index := columnIndexMap[group[i]]
		for _, items := range data {
			uniqueMap[items[index]] = struct{}{}
		}
		column := make([]string, 0, len(uniqueMap))
		for k := range uniqueMap {
			column = append(column, k)
		}
		sort.Slice(column, func(i, j int) bool {
			return column[i] < column[j]
		})
		tags[i] = tagInfo{values: column}
	}
	return tags
}

func filterData(columnIndexMap map[string]int, data [][]string, filters []PivotTableFilterOption) [][]string {
	if len(filters) == 0 || len(data) == 0 || len(columnIndexMap) == 0 {
		return data
	}
	var res [][]string
	for _, elem := range data {
		for _, filter := range filters {
			columnIndex, exist := columnIndexMap[filter.ColumnName]
			if !exist || columnIndex >= len(elem) {
				continue
			}
			if len(filter.Items) == 0 || isContains(filter.Items, elem[columnIndex]) {
				res = append(res, elem)
			}
		}
	}
	return res
}

func isContains(src []string, tar string) bool {
	for _, val := range src {
		if strings.EqualFold(tar, val) {
			return true
		}
	}
	return false
}
