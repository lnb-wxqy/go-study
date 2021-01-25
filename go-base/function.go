package go_base

import (
	"fmt"
	"math"
	"strings"
)

// 函数
// 函数变量（函数作为值）

// 1. 案例代码一
func String2Lower() {
	result := stringToLower("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
}

func processCase(str string) string {
	fmt.Println(str)
	result := ""
	// 奇数偶数依次显示为大小写
	for i, v := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(v))
		} else {
			result += strings.ToLower(string(v))
		}
	}
	return result
}

func stringToLower(s string, f func(string) string) string {
	fmt.Printf("%T\n", f)
	return f(s)
}

// 2. 案例2
type processFunc func(int) bool // 声明一个函数类型

func SliceOper() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice: ", slice)
	odd := filter(slice, isOdd) // 函数当做值传递
	fmt.Println("奇数元素：", odd)
	even := filter(slice, isEven)
	fmt.Println("偶数元素： ", even)

}

func filter(s []int, f processFunc) []int {
	sli := make([]int, 0)
	for _, v := range s {
		if f(v) {
			sli = append(sli, v)
		}
	}
	return sli
}

// 判断元素是否是奇数
func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

// 判断元素是否是偶数
func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

// 3. 匿名函数
// 3.1 匿名函数
func anonymousFunc() {
	func(data int) {
		fmt.Println("data: ", data)
	}(100)
}

// 3.2 将匿名函数赋值给变量
func anonymousFunc2() {
	f := func(data string) {
		fmt.Println(data)
	}
	f("life is left")
}

// 4. 匿名函数的用法-作回调函数
func callBackFunc() {
	arr := []float64{1, 9, 16, 25, 30}
	// 调用函数，对每个元素进行求平方根操作
	visit(arr, func(f float64) {
		f = math.Sqrt(f)
		fmt.Printf("%.2f\n", f)
	})
	// 调用函数，对每个元素进行求平方操作
	visit(arr, func(f float64) {
		f = math.Pow(f, 2)
		fmt.Printf("%.0f\n", f)
	})
}

// 定义一个函数，变量切片，对每个元素做处理
func visit(list []float64, f func(float64)) {
	for _, v := range list {
		f(v)
	}
}

// 5. 可变参数，参数可以传多个，也可以传切片slice
// 注意：a. 一个函数最多只能有一个可变参数
// 		 b. 参数类别中还有其他类型参数，则可变参数写在所有参数的最后
//语法：
//func 函数名(参数名 ...类型)(返回值列表){
//	// 函数体
//}
// 案例： 计算学员考试总成绩及平均成绩
func getScore(scores ...float64) (sum, avg float64, count int) {
	for _, s := range scores {
		sum += s
		count++
	}
	avg = sum / float64(count)
	return
}

func VariableArgFunc() {
	// 1. 传递n个参数
	sum, avg, count := getScore(90, 99.5, 88, 67, 89.78)
	fmt.Printf("学员公有%d门课，总成绩为：%.2f,平均成绩：%.2f\n", count, sum, avg)
	fmt.Println("-----------------------------------")

	// 2. 传递切片
	scores := []float64{90, 99.5, 88, 67, 89.78}
	sum, avg, count = getScore(scores...)
	fmt.Printf("学员公有%d门课，总成绩为：%.2f,平均成绩：%.2f\n", count, sum, avg)

}

// 6.递归函数
// 递归函数必须满足两个条件：
// a. 在每一次调用自己时，必须是（某种意义上）更接近于解
// b. 必须有一个终止处理或者计算的准则
// 案例：求阶乘
//计算阶乘n! = 1 x 2 x 3 x ... x n，⽤用函数fact(n)表示，可以看出：fact(n) =
//n! = 1 x 2 x 3 x ... x (n-1) x n = (n-1)! x n = fact(n-1) x n。所以，fact(n)可以
//表示为n x fact(n-1)，只有n=1时需要特殊处理理。

func recursion() {
	//result := factorial(10)
	result2 := factorial(50) //
	//result3 := fact(10,1)
	//fmt.Println("result: ", result)
	fmt.Println("result2: ", result2)
	//fmt.Println("result3: ", result3)
	fmt.Println("maxInt64: ", math.MaxInt64)
	fmt.Println("minInt64: ", math.MinInt64)

}

//n=100时，超出int64取值范围，输出结果是0，但是不抛异常，贼贼
func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 尾递归化
//使用递归函数需要注意防止栈溢出。在计算机中，函数调用是通过栈（stack）这种数据结构实现的，每当进入一个函数调用，
// 栈就会加一层栈帧，每当函数返回，栈就会减一层栈帧。由于栈的大小不是无限的，所以，递归调用的次数过多，会导致栈溢出
//尾递归是指，在函数返回的时候，调用自身本身，并且，return语句不能包含表达式。这样，
// 编译器或者解释器就可以把尾递归做优化， 使递归本身无论调用多少次，都只占用一个栈帧，不会出现栈溢出的情况。

// 递归函数优化
func fact(num, product int64) int64 {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return product
	}

	return fact(num-1, num*product)
}

// 7. 函数的参数传递
// 7.1 值传递
// 值传递： 是指在调用函数时将实际参数赋值一份传递到函数中，这样在函数中如果对参数进行修改不会影响到原内容数据，默认
//          情况下，Go语言使用的是值传递,这样会造成内存浪费、时间开销、性能降低的情况，可以传递指针来解决实参拷贝的问题
// 7.2 引用传递
// 引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数进行修改，将会影响原内容数据。
// 例1. 函数传int类型参数
func intFunc() {
	a := 10
	fmt.Printf("1. 变量a的内存地址：%p, 值为：%v\n\n", &a, a)
	// 传值
	changeIntVal(a)
	fmt.Printf("2、changeIntVal函数调⽤用之后：变量量a的内存地址：%p ，值为：%v \n\n", &a, a)

	//	 传引用
	changeIntPtr(&a)
	fmt.Printf("3、changeIntPtr函数调⽤用之后：变量量a的内存地址：%p ，值为：%v \n\n", &a, a) //50
}

func changeIntVal(a int) {
	fmt.Printf("--------changeIntVal函数内：值参数a的内存地址：%p ，值为：%v \n", &a, a) //10
	a = 90
}

func changeIntPtr(a *int) {
	fmt.Printf("--------changeIntPtr函数内：指针参数a的内存地址：%p ，值为：%v \n", &a, a) // 地址
	*a = 50
}
