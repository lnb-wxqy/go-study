package main

import "fmt"

var count int32

func main() {

	//atomic.AddInt32(&count,12)
	//
	//loadInt32 := atomic.LoadInt32(&count)
	//fmt.Println(count)
	//fmt.Println(loadInt32)

	// defer

	a := 1
	b := 2
	defer calc(a, calc(a, b))
	a = 0
	defer calc(a, calc(a, b))

	//	运行结果
	//	1 2 3
	//	0 2 2
	//	0 2 2
	//	1 3 4

	//slice
	s := []int{5}
	fmt.Println("cap: ", cap(s)) //1

	s = append(s, 7)
	fmt.Println("cap: ", cap(s)) //2

	s = append(s, 9)
	fmt.Println("cap: ", cap(s)) //4

	x := append(s, 11)
	fmt.Println(x)
	y := append(s, 12)
	fmt.Println(y)

	fmt.Println(s)
	fmt.Println(x)
	fmt.Println(y)

}

func calc(x, y int) int {
	fmt.Println(x, y, x+y)
	return x + y
}
