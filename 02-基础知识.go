/*
* @Author: csxiaoyao
* @Date:   2017-04-23 11:27:11
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-23 14:45:41
 */
/**
 * Go内置关键字和注释方法、Go程序的一般结构、包的导入、package别名与省略调用、可见性规则
 */

// 当前程序包名
package main

// 导包
import "fmt"

// 多包导入
import (
	"io"
	"os"
	"strings"
	"time"
)

// 别名
import std "fmt"

// 省略调用
import . "fmt"

// 常量定义
const PI = 3.14
const (
	PI1 = "3.14"
	PI2 = 3.14
)

// 全局变量声明与赋值
var name = "csxiaoyao"
var (
	name1 = 7
	name2 = "csxiaoyao"
)

// 一般类型声明
type newType int
type (
	type1 float32
	type2 byte
)

// 结构声明
type cs struct{}

// 接口声明
type sun interface{}

// 程序入口
func main() {

	/*
	   【Go内置关键字】（25个均为小写）
	   break	default			func 		interface	select
	   case	defer			go			map			struct
	   chan	else			goto		package		switch
	   const	fallthrough		if 			range		type
	   for		continue		import		return		var

	   【Go注释方法】
	   单行注释、多行注释

	   【Go程序的一般结构】：basic_structure.go
	   只有 package main（有且只有一个） 可包含 main 函数
	   import	导入非 main 包
	   const	常量定义
	   var		函数体外部使用全局变量的声明与赋值
	   type 	结构(struct)或接口(interface)的声明
	   func	函数声明

	   【可见性规则】
	   使用 大小写 决定该 常量、变量、类型、接口、结构、函数 是否可被外部包调用
	   函数名首字母 小写 即为 private
	   函数名首字母 大写 即为 public

	*/
	// fmt.Println(name)
	// std.Println(name)
	// Println(name)
}
