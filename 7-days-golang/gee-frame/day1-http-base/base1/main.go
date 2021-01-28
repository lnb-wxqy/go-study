package main

import (
	"fmt"
	"log"
	"net/http"
)

// go内置 net/http库简单使用

func main() {
	// 返回值
	// URL.Path = "/"
	http.HandleFunc("/", indexHandler)

	// 返回值：
	//Header["Connection"]=["keep-alive"]
	//Header["Upgrade-Insecure-Requests"]=["1"]
	//Header["User-Agent"]=["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"]
	//Header["Sec-Fetch-Mode"]=["navigate"]
	//Header["Accept-Encoding"]=["gzip, deflate, br"]
	//Header["Accept-Language"]=["zh,zh-CN;q=0.9"]
	//Header["Sec-Fetch-Dest"]=["document"]
	//Header["Accept"]=["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"]
	//Header["Sec-Fetch-Site"]=["none"]
	//Header["Sec-Fetch-User"]=["?1"]
	http.HandleFunc("/hello", helloHandler)

	// 第二个参数nil代表处理所有的http请求的实例，nil代表实用标准库中的实例处理
	// 第二个参数的类型是Handler，查看源码可知
	log.Fatalln(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}
