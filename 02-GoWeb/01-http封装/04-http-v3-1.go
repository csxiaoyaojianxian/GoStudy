/*
* @Author: csxiaoyao
* @Date:   2017-05-11 00:34:25
* @Last Modified by:   sunshine
* @Last Modified time: 2017-05-11 10:57:18
 */

package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// 自定义路由转发
var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		// 自动调用 hello 和 bye
		h(w, r)
		return
	}
	io.WriteString(w, "URL:"+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 3.")
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye, this is version 3.")
}
