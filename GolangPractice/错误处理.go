package main

import (
	"fmt"
)

// DivideError 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func main() {

	// 正常情况
	var a, b float64
	//var c string
	for true {

		fmt.Scanf("%f %f", &a, &b)

		if result, errorMsg := Divide(int(a), int(b)); errorMsg == "" {
			fmt.Println("100/100 = ", result)
		}
		// 当除数为零的时候会返回错误信息
		if _, errorMsg := Divide(int(a), int(b)); errorMsg != "" {
			fmt.Println("errorMsg is: ", errorMsg)
		}

		//if c == "exit" {
		//	os.Exit(0)
		//}

	}
}
