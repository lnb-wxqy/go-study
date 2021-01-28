package main

import (
	"go-study/7-days-golang/gee-frame/day3-router/geeb"
	"net/http"
)

func main() {
	r := geeb.New()

	r.Get("/", func(c *geeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Geeb!</h1>")
	})

	// 注：这个方法放到/hello前面，调用看着是正常的；
	// 但是如果放到/hello 方法后面，则调用 /19、/hello、/hello/name?name都会输出Hello 18！
	//r.Get("/18", func(c *geeb.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello 18!</h1>")
	//})

	r.Get("/hello", func(c *geeb.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
	})


	r.Get("/hello/:name", func(c *geeb.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
	})

	r.Get("/assets/*filepath", func(c *geeb.Context) {
		c.JSON(http.StatusOK, geeb.H{"filepath": c.Params["filepath"]})
	})

	r.Run(":9999")
}
