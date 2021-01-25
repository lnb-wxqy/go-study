package items

import "fmt"

type Demo struct{}

func (d *Demo) test() {
	fmt.Println("d: ", d)
	fmt.Println(d == nil)
}

func (d Demo) test2() {
	fmt.Println("d: ", d)
}

// d.test2()抛异常，应该是跟类型有关系
func MainFunc() {
	var d *Demo
	d.test()  // 正常调用
	d.test2() // panic: runtime error: invalid memory address or nil pointer dereference
}

func MainFunc2() {
	var d Demo
	d.test()  // 正常调用
	d.test2() // 正常调用
}
