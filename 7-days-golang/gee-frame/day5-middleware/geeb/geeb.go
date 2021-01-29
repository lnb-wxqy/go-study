package geeb

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type RouteGroup struct {
	prefix      string
	middlewares []HandlerFunc //support middleware
	parent      *RouteGroup   // support nesting
	engine      *Engine       // all groups share a Engine instance
}

type Engine struct {
	*RouteGroup
	router *router
	groups []*RouteGroup // store all groups
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

// Group is defined to create a new RouteGroup
// remember all groups share the same Engine instance
func (group *RouteGroup) Group(prefix string) *RouteGroup {
	engine := group.engine
	newGroup := &RouteGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouteGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

//可以仔细观察下addRoute函数，调用了group.engine.router.addRoute来实现了路由的映射。
//由于Engine从某种意义上继承了RouterGroup的所有属性和方法，
//因为 (*Engine).engine 是指向自己的。这样实现，我们既可以像原来一样添加路由，也可以通过分组添加路由。
func (group *RouteGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouteGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouteGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// run方法启动服务，http访问接口后会进到改方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
