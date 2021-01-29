### 第一天

1.http基础，简单介绍net/http库及http.Handler接口
2.搭建Geeb框架雏形。

### 第二天 

#### 一、上下文Context

1. 将路由(router)独立出来，方便之后增强
2. 设计上下文(Context)，封装Request和Response，提供对JOSN、HTML等返回类型的支持。
3. 设计Context分析：

- a. 对web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter。但是两个对象提供的接口粒度太细，比如我们要构造一个完整的响应，需要考虑消息头(Header)和消息体（Body),而Header包含了状态码(StatusCode)，消息类型(ContentType)等几乎每次请求都需要设置的信息。因此，如果不进行有效的封装，那么框架的用户将需要写大量重复，繁杂的代码，而且容易出错。针对常用场景，能够高效地构造出 HTTP 响应是一个好的框架必须考虑的点。

用返回 JSON 数据作比较，感受下封装前后的差距。
```go 
封装前：
obj = map[string]interface{}{
    "name": "geektutu",
    "password": "1234",
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
encoder := json.NewEncoder(w)
if err := encoder.Encode(obj); err != nil {
    http.Error(w, err.Error(), 500)
}
```

```go 
封装后：
c.JSON(http.StatusOK, gee.H{
    "username": c.PostForm("username"),
    "password": c.PostForm("password"),
})
```
- b.针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。


#### 二、 路由(Router)

- 我们将和路由相关的方法和结构提取了出来，放到了一个新的文件中router.go，方便我们下一次对 router 的功能进行增强，例如提供动态路由的支持。router 的 handle 方法作了一个细微的调整，即 handler 的参数，变成了 Context。


### 第三天

1. 使用Trie树实现动态路由（dynamic route）解析
2. 支持两种模式：name和 *filepath

#### 一、Trie树简介

之前，我们用了一个非常简单的map结构存储了路由表，使用map存储键值对，索引非常高效，但是有一个弊端，键值对的存储的方式，只能用来索引静态路由。那如果我我们想支持类似于 /hello/:name 这样的动态路由怎么办呢？所谓动态路由，即由一条路由规则可以匹配某一类型而非某一条固定的路由。例如：/hello/:name ，可以匹配 /hello/geektutu、/hello/jack等。

动态路由有很多种实现方式，支持的规则、性能等有很大的差异。例如开源的路由实现gorouter支持在路由规则中嵌入正则表达式，例如：/p/[0-9A-Za-z]+ ，即路径中的参数仅匹配数字和字母；另一个开源实现httprouter就不支持正则表达式。著名的Web开源框架gin在早期的版本，并没有实现自己的路由，而是直接使用了httprouter，后来不知道什么原因，放弃了httprouter，自己实现了一个版本。

![img.png](img.png)

实现动态路由最常用的数据结构，被称为前缀树（Trie树）。

每一个节点的所有的子节点都拥有相同的前缀，这种结构非常适用于路由匹配，比如定义了如下规则：

- /:lang/doc
- /:lang/tutorial
- /:lang/intro
- /about
- /p/blog
- /p/related

用前缀树表示是这样的：
![img_1.png](img_1.png)

HTTP请求的路径恰好是由 / 分隔的多段构成的，因此，每一段可以作为前缀树的一个节点。我们通过数结构查询，如果中间某一层的节点都不满足条件，那么就说明没有匹配到的路由，查询结束。

接下来我们实现的动态路由具备以下两个功能：

- 参数匹配： 例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
- 通配*：例如：/static/*filepath，可以匹配 /static/fav.ico，也可以匹配 /static/js/jQuery.js，这种模式常用于静态服务器，能够递归的匹配子路径。

#### 二、 Trie树实现

#### 三、 Context与handle的变化
在 HandlerFunc 中，希望能够访问到解析的参数，因此，需要对 Context 对象增加一个属性和方法，来提供对路由参数的访问。我们将解析后的参数存储到Params中，通过c.Param("lang")的方式获取到对应的值

### 第四天 分组控制Group
- 实现路由分组控制（Route Group Control）

#### 分组的意义
分组控制（Group Control）是Web框架应提供的基础功能之一。所谓分组，是指路由的分组。如果没有路由分组，我们需要针对每一个路由进行控制。但是真实的业务场景中，往往某一组路由需要相似的处理。例如：
- 以/post开头的路由匿名可访问。
- 以/admin开头的路由需要鉴权。
- 以/api开头的路由是 RESTful 接口，可以对接第三方平台，需要三方平台鉴权

大部分情况下的路由分组，是以相同的前缀来区分的。因此，我们今天实现的分组控制也是以前缀来区分，并且支持分组的嵌套。例如：/post是一个分组，/post/a和/post/b可以是该分组下的子分组。作用在/post分组上的中间件(middleware)，也都会作用在子分组，子分组还可以应用自己特有的中间件。

中间件可以给框架提供无限的扩展能力，应用在分组上，可以使得分组控制的收益更为明显，而不是共享相同的路由前缀这么简单。例如/admin的分组，可以应用鉴权中间件；/分组应用日志中间件，/是默认的最顶层的分组，也就意味着给所有的路由，即整个框架增加了记录日志的能力。
提供扩展能力支持中间件的内容，我们将在下一节当中介绍。

 分组嵌套
一个 Group 对象需要具备哪些属性呢？首先是前缀(prefix)，比如/，或者/api；要支持分组嵌套，那么需要知道当前分组的父亲(parent)是谁；当然了，按照我们一开始的分析，中间件是应用在分组上的，那还需要存储应用在该分组上的中间件(middlewares)。还记得，我们之前调用函数(*Engine).addRoute()来映射所有的路由规则和 Handler 。如果Group对象需要直接映射路由规则的话，比如我们想在使用框架时，这么调用：

 ```go
     r := gee.New()
     v1 := r.Group("/v1")
     v1.GET("/", func(c *gee.Context) {
     c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
     })
 ```
那么Group对象，还需要有访问Router的能力，为了方便，我们可以在Group中，保存一个指针，指向Engine，整个框架的所有资源都是由Engine统一协调的，那么就可以通过Engine间接地访问各种接口了。

### 第五天 中间件

- 设计并实现Web框架的中间件（Middlewares）机制
- 实现通用的logger中间件，能够记录请求到响应所花费的时间

#### 中间件是什么
中间件(middlewares)，简单说，就是非业务的技术类组件。Web 框架本身不可能去理解所有的业务，因而不可能实现所有的功能。因此，框架需要有一个插口，允许用户自己定义功能，嵌入到框架中，仿佛这个功能是框架原生支持的一样。因此，对中间件而言，需要考虑2个比较关键的点：
- 插入点在哪？使用框架的人并不关心底层逻辑的具体实现，如果插入点太底层，中间件逻辑就会非常复杂；如果插入点离用户太近，那和用户直接定义一组函数，每次在Handler中手工调用没有多大的优势了。
- 中间件的输入是什么？中间件的输入，决定了扩展能力。暴露的参数太少，用户发挥空间有限。
那对于一个web框架而言，中间件应该设计成什么样呢？接下来的实现，基本参考了gin框架。
#### 中间件的设计
Gee 的中间件的定义与路由映射的 Handler 一致，处理的输入是Context对象。插入点是框架接收到请求初始化Context对象后，允许用户使用自己定义的中间件做一些额外的处理，例如记录日志等，以及对Context进行二次加工。另外通过调用(*Context).Next()函数，中间件可等待用户自己定义的 Handler处理结束后，再做一些额外的操作，例如计算本次处理所用时间等。即 Gee 的中间件支持用户在请求被处理的前后，做一些额外的操作。举个例子，我们希望最终能够支持如下定义的中间件，c.Next()表示等待执行其他的中间件或用户的Handler：
```go
func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
```
另外，支持设置多个中间件，依次进行调用。

我们上一篇文章分组控制中讲到，中间件是应用在RouteGroup上的，应用在最顶层的Group，相当于作用于全局，所有的请求都会被中间件处理。那为什么不作用在每一条路由规则上呢？作用在某条路由规则，那还不如用户直接在Handler中调用直观。只作用在某条路由规则的功能通用性太差，不适合定义为中间件。

我们之前的框架设计是这样的，当接收到请求后，匹配路由，该请求的所有信息都保存在Context中。中间件也不例外，接收到请求后，应查找所有应作用于该路由的中间件，保存在Context中，依次进行调用。为什么依次调用后，还需要存在Context中呢？因为在设计中，中间件不仅作用在处理流程前，也可以作用在处理流程后，即在用户定义的Handler处理完毕后，还可以执行剩下的操作。

为此，我们给Context添加了两个参数，定义了Next方法：

```go
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path:   req.URL.Path,
		Method: req.Method,
		Req:    req,
		Writer: w,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}
```

index 是记录当前执行到第几个中间件，当在中间件中调用Next方法是，控制权交给了下一个中间件，然后在从后往前，调用每个中间件在Next方法之后定义的部分。如果我们将用户在映射路由时定义的Handler添加到c.handlers列表中，结果会怎么样呢？

```go
func A(c *Context) {
    part1
    c.Next()
    part2
}
func B(c *Context) {
    part3
    c.Next()
    part4
}
```

假设我们应用了中间件 A 和 B，和路由映射的 Handler。`c.handlers`是这样的[A, B, Handler]，`c.index`初始化为-1。调用`c.Next()`，接下来的流程是这样的：

- c.index++，c.index 变为 0
- 0 < 3，调用 c.handlers[0]，即 A
- 执行 part1，调用 c.Next()
- c.index++，c.index 变为 1
- 1 < 3，调用 c.handlers[1]，即 B
- 执行 part3，调用 c.Next()
- c.index++，c.index 变为 2
- 2 < 3，调用 c.handlers[2]，即Handler
- Handler 调用完毕，返回到 B 中的 part4，执行 part4
- part4 执行完毕，返回到 A 中的 part2，执行 part2
- part2 执行完毕，结束。

一句话说清楚重点，最终的顺序是`part1 -> part3 -> Handler -> part 4 -> part2`。恰恰满足了我们对中间件的要求，接下来看调用部分的代码，就能全部串起来了。

#### 代码实现

- 定义Use函数，将中间件应用到某个Group

``` go
// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

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
```

ServeHTTP 函数也有变化，当我们接收到一个具体请求时，要判断该请求适用于哪些中间件，在这里我们简单通过 URL 的前缀来判断。得到中间件列表后，赋值给 `c.handlers`。

- handle 函数中，将从路由匹配得到的 Handler 添加到 `c.handlers`列表中，执行`c.Next()`。

```go
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
```

#### 使用Demo

```go
func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
```

`gee.Logger()`即我们一开始就介绍的中间件，我们将这个中间件和框架代码放在了一起，作为框架默认提供的中间件。在这个例子中，我们将`gee.Logger()`应用在了全局，所有的路由都会应用该中间件。`onlyForV2()`是用来测试功能的，仅在`v2`对应的 Group 中应用了。

接下来使用 curl 测试，可以看到，v2 Group 2个中间件都生效了。

```
$ curl http://localhost:9999/
>>> log
2019/08/17 01:37:38 [200] / in 3.14µs

(2) global + group middleware
$ curl http://localhost:9999/v2/hello/geektutu
>>> log
2019/08/17 01:38:48 [200] /v2/hello/geektutu in 61.467µs for group v2
2019/08/17 01:38:48 [200] /v2/hello/geektutu in 281µs
```

### 第六天  模板（HTML Template）

- 实现静态资源服务（Static Resource）
- 支持HTML模板渲染

#### 服务端渲染

现在越来越流行前后端分离的开发模式，即Web后端提供RESTful接口，返回结构化的数据（通常为json或者xml）。前后端使用AJAX技术请求到所需的数据，利用JavaScript进行渲染。Vue/React等前端框架持续火热，这种开发模式前后端解耦，优势非常突出。后端童鞋专心解决资源利用，并发，数据库等问题，只需要考虑数据如何生成；前端童鞋专注于界面设计实现，只需要考虑拿到数据后如何渲染即可。使用 JSP 写过网站的童鞋，应该能感受到前后端耦合的痛苦。JSP 的表现力肯定是远不如 Vue/React 等专业做前端渲染的框架的。而且前后端分离在当前还有另外一个不可忽视的优势。因为后端只关注于数据，接口返回值是结构化的，与前端解耦。同一套后端服务能够同时支撑小程序、移动APP、PC端 Web 页面，以及对外提供的接口。随着前端工程化的不断地发展，Webpack，gulp 等工具层出不穷，前端技术越来越自成体系了。

但前后分离的一大问题在于，页面是在客户端渲染的，比如浏览器，这对于爬虫并不友好。Google 爬虫已经能够爬取渲染后的网页，但是短期内爬取服务端直接渲染的 HTML 页面仍是主流。

今天的内容便是介绍 Web 框架如何支持服务端渲染的场景。

#### 静态文件（Serve Static Files）

网页的三剑客，JavaScript、CSS 和 HTML。要做到服务端渲染，第一步便是要支持 JS、CSS 等静态文件。还记得我们之前设计动态路由的时候，支持通配符`*`匹配多级子路径。比如路由规则`/assets/*filepath`，可以匹配`/assets/`开头的所有的地址。例如`/assets/js/geektutu.js`，匹配后，参数`filepath`就赋值为`js/geektutu.js`。

那如果我们将所有的静态文件放在`/usr/web`目录下，那么`filepath`的值即是该目录下文件的相对地址。映射到真实的文件后，将文件返回，静态服务器就实现了。

找到文件后，如何返回这一步，`net/http`库已经实现了。因此，gee 框架要做的，仅仅是解析请求的地址，映射到服务器上文件的真实地址，交给`http.FileServer`处理就好了。

```go
// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}
```

我们给RouteGroup添加了2个方法，Static这个方法是暴露给用户的。用户可以将磁盘上的某个文件夹root映射到路由relativepath。例如：

```go
r := gee.New()
r.Static("/assets", "/usr/geektutu/blog/static")
// 或相对路径 r.Static("/assets", "./static")
r.Run(":9999")
```

用户访问 localhost:9999/asserts/js/geektutu.js，最终返回：localhost:9999/usr/geektutu/blog/static/js/geektutu.js

#### HTML模板渲染

Go语言内置了`text/template`和`html/template`2个模板标准库，其中[html/template](https://golang.org/pkg/html/template/)为 HTML 提供了较为完整的支持。包括普通变量渲染、列表渲染、对象渲染等。gee 框架的模板渲染直接使用了`html/template`提供的能力。

```go
Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup     // store all groups
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
```

首先为Engine示例添加了*template.Template和template.FuncMap对象，前者将所有的模板加载进内存，后者是所有的自定义模板渲染函数。

另外，给用户分别提供了设置自定义渲染函数funcMap和加载模板的方法。

接下来，对原来的(*Context).HTML()方法做了写下修改，使之支持根据模板文件名选择模板进行渲染。