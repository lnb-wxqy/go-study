第一天
1.http基础，简单介绍net/http库及http.Handler接口
2.搭建Geeb框架雏形。


第二天 
一、上下文Context
1. 将路由(router)独立出来，方便之后增强
2. 设计上下文(Context)，封装Request和Response，提供对JOSN、HTML等返回类型的支持。
3. 设计Context分析：
   3.1 必要性
   a. 对web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter。但是两个对象提供的接口粒度太细，比如
      我们要构造一个完整的响应，需要考虑消息头(Header)和消息体（Body),而Header包含了状态码(StatusCode)，消息类型(ContentType)
      等几乎每次请求都需要设置的信息。因此，如果不进行有效的封装，那么框架的用户将需要写大量重复，繁杂的代码，而且容易出错。针对常用场景，
      能够高效地构造出 HTTP 响应是一个好的框架必须考虑的点。

用返回 JSON 数据作比较，感受下封装前后的差距。

```go
//封装前：
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
//VS 封装后：

c.JSON(http.StatusOK, gee.H{
    "username": c.PostForm("username"),
    "password": c.PostForm("password"),
})
```
b.针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。
  对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？再比如，框架需要支持中间件，
  那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。
  因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用
  Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。

二、 路由(Router)
我们将和路由相关的方法和结构提取了出来，放到了一个新的文件中router.go，方便我们下一次对 router 的功能进行增强，
例如提供动态路由的支持。router 的 handle 方法作了一个细微的调整，即 handler 的参数，变成了 Context。