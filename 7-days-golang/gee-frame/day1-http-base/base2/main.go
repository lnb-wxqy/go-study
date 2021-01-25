package main

import (
	"fmt"
	"log"
	"net/http"
)

// 源码如下
//package http
//
//type Handler interface {
//	ServeHTTP(w ResponseWriter, r *Request)
//}
//
//func ListenAndServe(address string, h Handler) error

// -------------------------------------------------------
// 通过查看net/http的源码可以发现，ListenAndServe方法的第二个参数Handler是一个接口，需要
// 实现方法ServeHTTP，也就是说，只要传入任何实现了Handler接口的实例，所有的HTTP请求，就都交
// 给了该实例处理了。

// 定义一个结构体
type Engine struct {
}

// 实现Handler接口的方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/engine":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=[%q]\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)

	}
}

func main() {
	e := new(Engine)
	log.Fatal(http.ListenAndServe(":9998", e))
}
