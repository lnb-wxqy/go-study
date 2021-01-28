package geeb

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// run方法启动服务，http访问接口后会进到改方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("=================")
	c := newContext(w, req)
	engine.router.handle(c)
}
