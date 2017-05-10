/*
* @Author: csxiaoyao
* @Date:   2017-05-10 17:18:39
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 22:45:02
 */

package main

import "fmt"

/**
 * 【part1】
 */
// 定义一个 DivideError 结构
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

/**
 * 【part2】
 */
// 通过实现 error 接口类型来生成错误信息。函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息
// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, errors.New("math: square root of negative number")
// 	}
// 	// 实现
// 	return 0, nil
// }

func main() {
	// 通过内置的错误接口提供错误处理机制
	// error类型是一个接口类型
	// 【定义】：
	// type error interface {
	//     Error() string
	// }
	// 【实现】：
	// result, err := Sqrt(-1)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result) // 100/10 =  10
	}
	// 当被除数为零会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
	/*errorMsg is:
	Cannot proceed, the divider is zero.
	dividee: 100
	divider: 0*/
}
