package geeb

import (
	"fmt"
	"net/http"
)

// HandlerFunc
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine struct,implements ServeHTTP function
type Engine struct {
	router map[string]HandlerFunc
}

// Engine new function
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc, 0)}
}

// addRouter
func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET
func (engine *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", pattern, handlerFunc)
}

// POST
func (engine *Engine) POSt(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", pattern, handlerFunc)
}

// Run
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// implements ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
