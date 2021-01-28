package geeb

import (
	"net/http"
	"strings"
)

//Trie 树的插入与查找都成功实现了，接下来我们将 Trie 树应用到路由中去吧。我们使用 roots 来存储每种请求方式的Trie 树根节点。
//使用 handlers 存储每种请求方式的 HandlerFunc 。getRoute 函数中，还解析了:和*两种匹配符的参数，返回一个 map 。
//例如/p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}，/static/css/geektutu.css匹配到/static/*filepath，
//解析结果为{filepath: "css/geektutu.css"}。

// roots key eg,roots['GET'] roots['POST']
// handlers key eg,handlers['GET-/p/:lang/doc'],handlers['POST-/p/book']

type router struct {
	roots    map[string]*node // roots用来存储每种请求方式的Trie树根节点
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node, 0),
		handlers: make(map[string]HandlerFunc, 0),
	}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string, 0)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	node := root.search(searchParts, 0)
	if node == nil {
		return nil, nil
	}

	parts := parsePattern(node.pattern)
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}
	return node, params

}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes

}

func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
