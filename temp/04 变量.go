/*
* @Author: sunshine
* @Date:   2017-04-22 20:38:06
* @Last Modified by:   sunshine
* @Last Modified time: 2017-04-22 21:15:38
*/

package main

import "fmt"

func main() {
// 【变量声明】
// (1) 指定变量类型，声明后若不赋值，使用默认值
// var v_name v_type
// v_name = value
// (2) 根据值自行判定变量类型。
// var v_name = value
// (3) 省略var, 注意 := 左侧的变量不该是已经声明过的，否则会编译错误
// v_name := value

var a int = 10
var b = 10
c : = 10

// 【多变量声明】
// 类型相同多个变量, 非全局变量
// var vname1, vname2, vname3 type
// vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 //和python很像,不需要显式声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 //出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)

}