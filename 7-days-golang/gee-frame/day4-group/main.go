package main

import (
	"go-study/7-days-golang/gee-frame/day4-group/geeb"
	"net/http"
)

func main() {
	r := geeb.New()
	r.GET("/index", func(c *geeb.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	v1.GET("/", func(c *geeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Geeb</h1>")
	})
	v1.GET("/hello", func(c *geeb.Context) {
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
	})

	v2 := r.Group("/v2")
	v2.GET("/hello/:name", func(c *geeb.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	v2.POST("/login", func(c *geeb.Context) {
		c.JSON(http.StatusOK, geeb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
