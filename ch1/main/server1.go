package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 每个请求都调用handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 回显请求的path
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.RequestURI())
}