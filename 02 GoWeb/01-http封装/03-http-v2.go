/*
* @Author: csxiaoyao
* @Date:   2017-05-11 00:34:25
* @Last Modified by:   sunshine
* @Last Modified time: 2017-05-11 10:27:31
 */

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL:"+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 2.")
}
