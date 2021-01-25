package go_base

import "fmt"

// 位运算符
//按位与（&）: 同一位同时为1则为1，反之为 0
//按位或（|）：同一位其中一个为1则为1，反之为 0
//按位异或（^）：同一位不同时为1，反之为 0
//按位左移（<<）：左移1位相当于乘以2的1次方，左移n位相当于乘以2的n次方
//按位右移（>>）：右移1位相当于除以2的1次方，右移n位相当于除以2的n次方
func BitwiseOperator() {
	var a uint = 60 // 60 = 0011 1100
	var b uint = 13 // 13 = 0000 1101
	var c uint = 0

	c = a & b
	fmt.Printf("a&b的结果c的值为：%d\n", c) // 结果：12

	c = a | b
	fmt.Printf("a|b的结果c的值为：%d\n", c) // 结果：61

	c = a ^ b
	fmt.Printf("a^b的结果c的值为：%d\n", c) // 结果：49

	c = a << 2
	fmt.Printf("a<<2的结果c的值为：%d\n", c) // 结果：240

	c = a >> 2
	fmt.Printf("a>>b的结果c的值为：%d\n", c) // 结果：15
}
