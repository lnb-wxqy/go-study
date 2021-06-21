package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
	"net/http"
)

// 爬取简书首页文章列表

//回调函数的执行顺序
//1. OnRequest 请求发出之前调用
//2. OnError 请求过程中出现Error时调用
//3. OnResponse 收到response后调用
//4. OnHTML 如果收到的内容是HTML，就在onResponse执行后调用
//5. OnXML 如果收到的内容是HTML或者XML，就在onHTML执行后调用
//6. OnScraped OnXML执行后调用

var jsurl = "https://www.jianshu.com"

var cookie2 = `__gads=ID=a493635b6240cc77:T=1580533301:S=ALNI_MaURrztW_z1xmPG3NKdIC1vWwDWVg; _ga=GA1.2.1821575570.1587090178; __yadk_uid=9UoH3oJ4mDAP7VeM8q3NVcR1rknVjYxu; UM_distinctid=178c42d6b3c834-0988df9ba636a6-f313f6d-e1000-178c42d6b3d504; read_mode=night; default_font=font2; locale=zh-CN; _gid=GA1.2.909701685.1622708318; CNZZDATA1279807957=169090193-1618198732-https%3A%2F%2Fwww.baidu.com%2F|1622709064; sensorsdata2015jssdkcross={"distinct_id":"175bb400f7e4df-036c79ff5331ff-f313f6d-921600-175bb400f7f634","first_id":"","props":{"$latest_traffic_source_type":"社交网站流量","$latest_search_keyword":"未取到值","$latest_referrer":"https://open.weixin.qq.com/","$latest_utm_source":"desktop","$latest_utm_medium":"timeline","$latest_utm_campaign":"maleskine","$latest_utm_content":"note"},"$device_id":"175bb400f7e4df-036c79ff5331ff-f313f6d-921600-175bb400f7f634"}; Hm_lvt_0c0e9d9b1e7d617b3e6842e85b9fb068=1622431485,1622450831,1622708318,1622711021; Hm_lpvt_0c0e9d9b1e7d617b3e6842e85b9fb068=1622711113; remember_user_token=W1sxNTc2NDgwOF0sIiQyYSQxMSR0UVp4YkRxNDkzazk4MTFpcjd1SXBlIiwiMTYyMjcxMTE2NS43NTg3NTE0Il0=--b9d63ba368874e2668f78f6196167a57243cf223; web_login_version=MTYyMjcxMTE2NQ==--85c468eba0f7a135c9bc94302bfd6290263d28ff; _m7e_session_core=b923c9f4f231f210015a453637ca08a5; _gat=1`

func main() {
	// 创建采集器1
	c1 := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		colly.MaxDepth(1),
		colly.Debugger(&debug.LogDebugger{}))

	// 采集器1，获取文章列表
	c1.OnHTML("ul[class= 'note-list']", func(e *colly.HTMLElement) {
		//列表中的每一项
		e.ForEach("li", func(i int, element *colly.HTMLElement) {
			//  文章链接
			href := element.ChildAttr("div[class= 'content'] > a[class='title']", "href")
			// 文章标题
			title := element.ChildText("div[class= 'content'] > a[class='title']")

			//文章摘要
			summary := element.ChildText("div[class='content'] > p[class='abstract']")

			fmt.Println(title, href)
			fmt.Println(summary)
			fmt.Println()

		})
	})

	// 爬取需要登录的网页
	//官网提供登录页处理的例子，但是大多数涉及验证码，不好处理，目前方式是手动登录，复制cookie写到爬虫请求头里
	extensions.Referer(c1)
	// 设置登录cookie
	c1.SetCookies(jsurl, []*http.Cookie{
		&http.Cookie{
			Name:     "remember_user_token",
			Value:    cookie2,
			Path:     "/",
			Domain:   ".jianshu.com",
			Secure:   true,
			HttpOnly: true,
		},
	})

	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("1爬取页面： ", r.URL)
	})

	c1.OnError(func(response *colly.Response, err error) {
		fmt.Println("Request URL: ", response.Request.URL, "failed with response: ", response, "\nError", err)
	})

	// 访问具体网站
	err := c1.Visit("https://www.jianshu.com/")
	if err != nil {
		fmt.Println("访问完整地址err: ", err)
	}

}
