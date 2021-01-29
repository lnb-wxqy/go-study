package geeb

import (
	"log"
	"time"
)

//Gee 的中间件的定义与路由映射的 Handler 一致，处理的输入是Context对象。
//插入点是框架接收到请求初始化Context对象后，允许用户使用自己定义的中间件做
//一些额外的处理，例如记录日志等，以及对Context进行二次加工。另外通过调
//用(*Context).Next()函数，中间件可等待用户自己定义的 Handler处理结束后，
//再做一些额外的操作，例如计算本次处理所用时间等。即 Gee 的中间件支持用户在请
//求被处理的前后，做一些额外的操作。举个例子，我们希望最终能够支持如下定义的中间件，
//c.Next()表示等待执行其他的中间件或用户的Handler:

func Logger() HandlerFunc {
	return func(c *Context) {
		// start time
		t := time.Now()
		// process reques
		c.Next()
		//calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
