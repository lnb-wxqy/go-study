package main

import (
	"go-study/7-days-golang/gee-frame/day2-context/geeb"
	"net/http"
)

func main() {
	engine := geeb.New()

	engine.Get("/", func(c *geeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Geeb <h1>")
	})

	engine.Get("/hello", func(c *geeb.Context) {
		c.String(http.StatusOK, "hello %s ,you are at %s\n", c.Query("name"), c.Path)
	})

	// 127.0.0.1:9884/login?username=geektutu&password=123466
	engine.Post("/login", func(c *geeb.Context) {
		c.JSON(http.StatusOK, geeb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	engine.Run(":9884")
}
