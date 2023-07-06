package main

func main() {
	//s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"
	//// call top level function
	//r, err := engine.ParseAndExec(s)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("%s = %v\n", s, r)

	//s := "@ Variable A 1 - @ Variable B + @ Variable C"
	//_, err := engine.Parse(s)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//exp := NewExpression(s)
	//exp.AddParam("@ Variable A", 1)
	//exp.AddParam("@ Variable B", 2.23)
	//exp.AddParam("@ Variable C", 3)
	//res, err := exp.Exec()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res)
}

//type Expression struct {
//	expression string
//	params     map[string]any
//}
//
//func NewExpression(exp string) (*Expression, error) {
//	return &Expression{
//		expression: exp,
//		params:     map[string]any{},
//	}, nil
//}
//
//func (e *Expression) AddParam(variable string, value any) {
//	e.params[variable] = value
//}
//
//func (e *Expression) Exec() (float64, error) {
//	exp := e.expression
//	for k, v := range e.params {
//		exp = strings.ReplaceAll(exp, k, fmt.Sprintf("%v", v))
//	}
//	res, err := engine.ParseAndExec(exp)
//	if err != nil {
//		return 0, fmt.Errorf("公式格式错误")
//	}
//	return res, nil
//}
