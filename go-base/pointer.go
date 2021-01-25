package go_base

import "fmt"

// 指针： 指针是存储另一个变量的内存地址的变量
// 特点： 指针不能运算
// 示例1. 获取变量的地址，使用&
func GetVAddress() {
	a := 10
	fmt.Printf("a变量的地址：%x\n", &a)

}

//示例2
// 运行结果：
//a变量的地址：c00000a368
//a的类型：int, 值：120
//&a的类型：*int, 值：0xc00000a370
//*&a的类型：int, 值：120
//p的类型：*int, 值：0xc00000a370
//*p的类型：int, 值：120
//120 0xc00000a370 120
//0xc00000a370 0xc000006038 120 0xc00000a370 0xc00000a370
func pointerFunc() {
	// 声明实际变量
	var a int = 120

	// 声明指针 变量
	var p *int

	// 给指针变量赋值，将变a的地址赋值给p
	p = &a

	fmt.Printf("a的类型：%T, 值：%v\n", a, a)
	fmt.Printf("&a的类型：%T, 值：%v\n", &a, &a)
	fmt.Printf("*&a的类型：%T, 值：%v\n", *&a, *&a)

	fmt.Printf("p的类型：%T, 值：%v\n", p, p)
	fmt.Printf("*p的类型：%T, 值：%v\n", *p, *p)

	fmt.Println(a, &a, *&a)
	fmt.Println(p, &p, *p, *(&p), &(*p))
}
