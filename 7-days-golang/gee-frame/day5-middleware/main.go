package main

import (
	"go-study/7-days-golang/gee-frame/day5-middleware/geeb"
	"log"
	"net/http"
	"time"
)

func onlyForV2() geeb.HandlerFunc {
	return func(c *geeb.Context) {
		// start time
		t := time.Now()

		// if a server error occurred
		c.Fail(500, "Internal Server Error")

		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r :=geeb.New()
	r.Use(geeb.Logger()) // global middleware
	r.GET("/", func(c *geeb.Context) {
		c.HTML(http.StatusOK,"<h1>Hello Geeb</h1>")
	})

	v2 :=r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware

	v2.GET("/hello/:name", func(c *geeb.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Run(":9999")
}