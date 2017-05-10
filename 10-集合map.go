/*
* @Author: csxiaoyao
* @Date:   2017-04-23 21:31:57
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-10 16:53:59
 */
/**
 * map 概述、简单 map 的创建与使用、复杂 map 与键值对操作、map 与 slice 的迭代操作、元素类型为 map 的 slice、map 的间接排序
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	   【map】
	   类似其它语言中的哈希表或者字典，以key-value形式存储数据，无序
	   Key必须支持==或!=比较运算类型，不可以是函数、map或slice
	   Map查找比线性搜索快很多，但比索引访问数据慢100倍
	   Map使用make()创建，支持 := 简写方式

	   make([keyType]valueType, cap)，cap表示容量，可省略，超出容量时会自动扩容
	   键值对不存在时自动添加，使用delete()删除某键值对
	   使用 for range 对map和slice进行迭代操作
	*/
	// 【初始化】
	//
	// var m map[int]string
	// m = map[int]string{}
	// fmt.Println(m)

	var m1 map[int]string = make(map[int]string)
	fmt.Println(m1) // map[]

	m2 := make(map[int]string)
	fmt.Println(m2) // map[]

	// 【存取】
	// 简单map
	m1[1] = "OK"
	fmt.Println(m1)    // map[1:OK]
	delete(m1, 1)      // 删除元素：map_name,key_name
	fmt.Println(m1[1]) //
	// 复杂map，多维map
	var m3 map[int]map[int]string
	m3 = make(map[int]map[int]string)

	// 【使用第二参数判断集合中元素是否存在】
	a, ok := m3[2][1]
	if !ok { // ok 为 true 代表存在
		m3[2] = make(map[int]string)
	}
	m3[2][1] = "GOOD"
	a, ok = m3[2][1]
	fmt.Println(a, ok) // GOOD true

	// 【迭代】
	sm := make([]map[int]string, 5)
	for key := range sm {
		sm[key] = make(map[int]string, 1)
		sm[key][1] = "OK"
		fmt.Println(sm[key])
	}
	fmt.Println(sm) // [map[1:OK] map[1:OK] map[1:OK] map[1:OK] map[1:OK]]

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)   // 间接排序
	fmt.Println(s) // [1 2 3 4 5]

	// 【对调键值】
	map1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(map1) // map[1:a 2:b 3:c 4:d 5:e]
	map2 := make(map[string]int)
	for k, v := range map1 {
		map2[v] = k
	}
	fmt.Println(map2) // map[b:2 c:3 d:4 e:5 a:1]
}
