/*
* @Author: csxiaoyao
* @Date:   2017-04-25 00:41:34
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-25 00:56:16
 */

package main

/**
 * 反射基本操作、反射匿名或嵌入字段、修改目标对象、动态调用方法
 */
import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World!")
}

func main() {
	u := User{1, "OK", 12}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println("%6s:%v = %vn", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("%6s:%vn", m.Name, m.Type)
	}
}

// 结果
// Type: User
// Fields:
// %6s:%v = %vn Id int 1
// %6s:%v = %vn Name string OK
// %6s:%v = %vn Age int 12
// %6s:%vn Hello func(main.User)
