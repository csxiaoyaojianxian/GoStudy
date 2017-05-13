/*
* @Author: csxiaoyao
* @Date:   2017-05-11 00:34:25
* @Last Modified by:   sunshine
* @Last Modified time: 2017-05-11 10:59:19
 */

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 路径处理
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL:"+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 3.")
}
