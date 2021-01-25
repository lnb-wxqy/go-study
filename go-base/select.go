package go_base

import "fmt"

// select语句类似于switch，但是select会随机执行一个可运行的case。
// 如果没有case可运行，它将阻塞，直到有case可运行。

func SelectCase() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, "from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, "to c2")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("received ", i3, "from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}
}

