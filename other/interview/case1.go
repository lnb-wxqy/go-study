package interview

import "fmt"

// 1.写出下面代码输出内容
//考点：defer执行顺序
//解答：defer是后进先出
//panic需要等defer结束后才会向上传递。出现panic时，会先按照
//defer的后入先出的顺序执行，最后才会执行panic
//

func Defer_call()  {
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()
	panic("触发异常")
}
