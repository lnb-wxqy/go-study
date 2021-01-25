package go_base

import (
	"fmt"
)

// Type switch
// switch 语句还可以被用于type-switch来判断某个interface变量中的实际类型
// 语法结构：
/*
	switch x.(type){
		case type:
			statements(s)
		case type:
			statements(s)
		default:
			statements(s)
	}
*/

// 代码示例
func SwitchType() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T\n", i)
	case int:
		fmt.Println("x 是int类型")
	case float64:
		fmt.Println("x 是 float64 类型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 类型")
	default:
		fmt.Println("未知类型")

	}
}
