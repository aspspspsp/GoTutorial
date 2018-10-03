package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func sayFuck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "a") //这个写入到w的是输出到客户端的
	fmt.Println("test")
}

func sayFuckYou(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "b") //这个写入到w的是输出到客户端的
	fmt.Println("test2")
}

func main() {
	//http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/", sayFuckYou)         //设置访问的路由
	http.HandleFunc("/fuck", sayFuck)        //设置访问的路由
	err := http.ListenAndServe(":9999", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
