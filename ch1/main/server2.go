package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 避免竞态条件
var mu sync.Mutex
// 数量count
var count int

func main() {
	// http.HandleFunc("/", handler2)
	http.HandleFunc("/", handler3)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 一个handler函数
func handler2(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count ++
	mu.Unlock()
	fmt.Fprintf(w, "URL.PATH: %q\n", r.URL.Path)
}

// handler将请求的http头和form数据都打印出来，这样可以使得调试更加的方便
func handler3(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	// 循环打印Header
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header[%q]\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil{
		log.Fatal(err)
	}
	for k, v := range r.Form{
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Counte: %d\n", count)
	mu.Unlock()
}