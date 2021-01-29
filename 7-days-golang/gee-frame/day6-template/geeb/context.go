package geeb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 设计Context

// define map H
type H map[string]interface{}

// context struct
type Context struct {
	// origin objs
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string

	//response info
	StatusCode int

	// middleware
	handlers []HandlerFunc
	index    int // index是记录当前执行到第几个中间件

	//engine pointer
	engine *Engine
}

// newContext function
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

//
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// Param
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm
//方法参数是这种形式： 127.0.0.1:9884/login?username=geektutu&password=123466
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// status
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// setHeader
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// string
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Context-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// json
// 127.0.0.1:9884/login?username=geektutu&password=123466
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Context-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

// data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// html
// refer https://golang.org/pkg/html/template/
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Context-Type", "text/html")
	c.Status(code)
	//c.Writer.Write([]byte(html))
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}
