package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 使用make创建了一个传递string类型的channel
	ch := make(chan string)
	// 对每一个命令行参数，都用go关键字来创建一个goroutine
	// 并且在函数fetch内异步执行http.Get方法
	for _, url := range os.Args[1:]{
		// 创建一个新的goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:]{
		// main函数接收fetch往channel里发送的值
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		// 发送到channel ch
		ch <- fmt.Sprint(err)
		return
	}
	// io.Copy会把响应的Body内容拷贝到Discard输出流中
	// 可以把这个变量看作垃圾桶，可以向里面写一些不需要的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
