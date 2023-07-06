package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/dengsgo/math-engine/engine"
	"testing"
)

func Test_2(t *testing.T) {
	s := "5-3+3*6/10.00"
	//s := "var1 + 2 * var2 / 4 + (456 - 8 * 9.2) > 0"
	ee, err := govaluate.NewEvaluableExpression(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	params := make(map[string]interface{})
	params["var1"] = 1
	params["var2"] = 6
	res, err := ee.Evaluate(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s = %v\n", s, res)
}

func Test_1(t *testing.T) {
	//s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"
	s := "1 + 2 * 5 / 2.50"
	r, err := engine.ParseAndExec(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s = %v\n", s, r)
}

//ee, err := govaluate.NewEvaluableExpression(expStr)
//if err != nil {
//	return 0, err
//}
//// 参数不支持一些字符，考虑下处理变量 todo wxm
//res, err := ee.Evaluate(nil)
//if err != nil {
//	return 0, code.Formula.InvalidFormula
//}
//val, yes := res.(float64)
//if yes {
//	return val, nil
//}
//return 0, code.Formula.UnmatchedResults.WithMsg(fmt.Sprintf("result type not float, [%v]", res), fmt.Sprintf("运算结果不是浮点类型, [%v]", res))

type coverageInfo struct {
	fileName   string
	total      int
	coverage   []int
	unCoverage []int
}

func Test(t *testing.T) {
	ciMap := make(map[string]coverageInfo)
	t.Log(ciMap["xxx"])
}
