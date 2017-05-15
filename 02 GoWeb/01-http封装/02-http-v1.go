/*
* @Author: csxiaoyao
* @Date:   2017-05-11 00:34:25
* @Last Modified by:   csxiaoyao
* @Last Modified time: 2017-05-11 00:59:20
 */

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 设置路由
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 1.")
}
