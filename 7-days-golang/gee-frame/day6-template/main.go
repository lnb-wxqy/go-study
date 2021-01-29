package main

import (
	"fmt"
	"go-study/7-days-golang/gee-frame/day6-template/geeb"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormAsData(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := geeb.New()
	r.Use(geeb.Logger())

	r.SetFuncMap(template.FuncMap{
		"FormAsData": FormAsData,
	})

	templatePath :="E:/lnb/go/src/go-study/7-days-golang/gee-frame/day6-template/templates/"
	r.LoadHTMLGlob(templatePath)
	r.Static("asserts", ".static")

	stu1 := &student{Name: "Jack", Age: 20}
	stu2 := &student{Name: "Rose", Age: 22}

	r.GET("/", func(c *geeb.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", geeb.H{
			"title":  "geeb",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *geeb.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", geeb.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
