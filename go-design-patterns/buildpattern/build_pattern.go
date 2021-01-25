package buildpackage

import "fmt"

// 队列
// 建造者模式使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式

// 电脑的制造
type Computer interface {
	// 主板制造
	MakeCpu()

	// 键盘制造
	MakeKeyBoard()

	// 屏幕制造
	MakeScreen()
}

// 定义一个构造器
type Creator struct {
	computer Computer
}

// 创建一个构造器
func (c *Creator) Construct() *Computer {
	c.computer.MakeCpu()
	c.computer.MakeKeyBoard()
	c.computer.MakeScreen()
	return &c.computer
}

type HuaWeiComputer struct {
	Cpu      string
	KeyBoard string
	Screen   string
}

func (h HuaWeiComputer) MakeCpu() {
	fmt.Println("CPU 制造中...")
	h.Cpu = "cpu"
}

func (h HuaWeiComputer) MakeKeyBoard() {
	fmt.Println("键盘制造中...")
	h.KeyBoard = "keyboard"
}

func (h HuaWeiComputer) MakeScreen() {
	fmt.Println("屏幕制造中...")
	h.Screen = "screen"
}

//func main() {
//	creator := Creator{
//		computer: HuaWeiComputer{},
//	}
//	creator.Construct()
//	fmt.Printf("%v", creator)
//}
