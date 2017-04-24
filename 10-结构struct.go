/*
* @Author: csxiaoyao
* @Date:   2017-04-24 11:16:23
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-04-24 14:46:36
 */
/**
 * 结构的定义与使用、使用字面值初始化、匿名结构与字段、结构间的赋值与比较、嵌入结构
 */
package main

import "fmt"

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}
type person2 struct {
	string
	int
}

// struct 嵌套
type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {

	/*
	   【结构struct】
	   默认值传递，支持指向自身的指针类型成员
	   支持匿名结构，可用作成员或定义成员变量
	   匿名结构也可以用于map的值
	   可以使用字面值对结构进行初始化
	   允许直接通过指针来读写结构成员
	   相同类型的成员可进行直接拷贝赋值
	   支持 == 与 !=比较运算符，但不支持 > 或 <
	   支持匿名字段，本质上是定义了以某个类型名为名称的字段
	   嵌入结构作为匿名字段看起来像继承，但不是继承
	   可以使用匿名字段指针


	*/
	a := person{
		Name: "sunshine",
		Age:  25,
	}
	a.Age = 24
	a.Contact.Phone = "11111111111"
	a.Contact.City = "sz"
	A(&a)
	fmt.Println(a) // {csxiaoyao 24 {11111111111 sz}}

	// 匿名
	b := person2{"sun", 25}
	fmt.Println(b) // {sun 25}

	// 嵌套
	c := student{Name: "csxiaoyao", Age: 25, human: human{Sex: 0}}
	fmt.Println(c) // {{0} csxiaoyao 25}
}
func A(per *person) {
	per.Name = "csxiaoyao"
}
