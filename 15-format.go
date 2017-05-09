/*
* @Author: csxiaoyao
* @Date:   2017-05-08 20:32:45
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-09 23:29:07
 */

package main

import "fmt"

type Human struct {
	Name string
}

var people = Human{Name: "sunshine"}

func main() {
	/**
	 * 【普通占位符】
	 */
	// 相应值的默认格式
	fmt.Printf("%v \n", people) // {sunshine}
	// 打印结构体时，会添加字段名
	fmt.Printf("%+v \n", people) // {Name:sunshine}
	// 相应值的Go语法表示
	fmt.Printf("%#v \n", people) // main.Human{Name:"sunshine"}
	// 相应值的类型的Go语法表示
	fmt.Printf("%T \n", people) // main.Human
	// 字面上的百分号，并非值的占位符
	fmt.Printf("%% \n", people) // %

	/**
	 * 【布尔占位符】
	 */
	// true 或 false
	fmt.Printf("%t \n", true) // true

	/**
	 * 【整数占位符】
	 */
	// 二进制表示
	fmt.Printf("%b \n", 5) // 101
	// 相应Unicode码点所表示的字符
	fmt.Printf("%c \n", 0x4E2D) // 中
	// 十进制表示
	fmt.Printf("%d \n", 0x12) // 18
	// 八进制表示
	fmt.Printf("%o \n", 10) // 12
	// 单引号围绕的字符字面值，由Go语法安全地转义
	fmt.Printf("%q \n", 0x4E2D) // '中'
	// 十六进制表示，字母形式为小写 a-f
	fmt.Printf("%x \n", 13) // d
	// 十六进制表示，字母形式为大写 A-F
	fmt.Printf("%X \n", 13) // D
	// Unicode格式：U+1234，等同于 "U+%04X"
	fmt.Printf("%U \n", '中') // U+4E2D

	/**
	 * 【浮点数和复数的组成部分（实部和虚部）】
	 */
	// 科学计数法
	fmt.Printf("%e , %E \n", 10.2, 10.2) // 1.020000e+01 , 1.020000E+01
	// 有小数点而无指数
	fmt.Printf("%f \n", 10.2) // 10.200000
	// 根据情况选择 `%e` 或 `%f` 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("%g \n", 10.20) // 10.2
	// 根据情况选择 `%E` 或 `%f` 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("%G \n", 10.20+2i) // (10.2+2i)

	/**
	 * 【字符串与字节切片】
	 */
	// 输出字符串表示（string类型或[]byte）
	fmt.Printf("%s \n", []byte("Go语言")) // Go语言
	// 双引号围绕的字符串，由Go语法安全地转义
	fmt.Printf("%q \n", "Go语言") // "Go语言"
	// 十六进制，小写字母，每字节两个字符
	fmt.Printf("%x \n", "golang") // 676f6c616e67
	// 十六进制，大写字母，每字节两个字符
	fmt.Printf("%X \n", "golang") // 676F6C616E67

	/**
	 * 【指针】
	 */
	// 十六进制表示，前缀 0x
	var people string
	fmt.Printf("%p \n", &people) // 0xc04200a370

	/**
	 * 【其它标记】
	 */
	// `+` 总打印数值的正负号，对于%q（%+q）保证只输出ASCII编码的字符
	fmt.Printf("%+d \n", 0xA)  // +10
	fmt.Printf("%+q \n", "中文") // "\u4e2d\u6587"
	// `-` 在右侧而非左侧填充空格（左对齐该区域）

	// `#` 备用格式：为八进制添加前导 0（%#o）
	fmt.Printf("%#U \n", '中') // U+4E2D '中'
	// 宽度与精度
	fmt.Printf("%.4g \n", 123.45)  // 123.5
	fmt.Printf("%6.2f \n", 123.45) // 123.45

}
