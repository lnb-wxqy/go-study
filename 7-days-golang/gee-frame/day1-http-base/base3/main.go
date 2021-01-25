package main

import (
	"fmt"
	"go-study/7-days-golang/gee-frame/day1-http-base/base3/geeb"
	"net/http"
)

func main() {
	r := geeb.New()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q]  %q\n", k, v)
		}
	})

	r.Run(":9886")
}
