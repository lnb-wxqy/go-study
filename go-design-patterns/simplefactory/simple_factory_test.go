package simplefactory

import "testing"

func TestOperationFactoryPattern(t *testing.T) {
	var ret int
	ret = OperationFactory("+").Exe(10, 5)
	t.Log("ret: ", ret)
	ret = OperationFactory("-").Exe(10, 5)
	t.Log("ret: ", ret)
	ret = OperationFactory("*").Exe(10, 5)
	t.Log("ret: ", ret)
	ret = OperationFactory("/").Exe(10, 5)
	t.Log("ret: ", ret)
}
