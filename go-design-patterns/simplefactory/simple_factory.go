package simplefactory

// 简单工厂模式
// 定义一个提供产品的接口，有子类决定生产什么，类似于java中的多态

type Operation interface {
	Exe(a, b int) int
}

type OperationAdd struct{}
type OperationSub struct{}
type OperationMul struct{}
type OperationDiv struct{}

func (operation *OperationAdd) Exe(a, b int) int {
	return a + b
}

func (operation *OperationSub) Exe(a, b int) int {
	return a - b
}

func (operation *OperationMul) Exe(a, b int) int {
	return a * b
}

func (operation *OperationDiv) Exe(a, b int) int {
	return a / b
}

func OperationFactory(oper string) Operation {
	switch oper {
	case "+":
		return &OperationAdd{}
	case "-":
		return &OperationSub{}
	case "*":
		return &OperationMul{}
	case "/":
		return &OperationDiv{}

	}
	return nil
}
