package reflect_lib

import "testing"

func TestTypeOfMethod(t *testing.T) {
	TypeOfMethod()
	t.Log("----------")
	TypeOfMethod2()
	ReflectStruct()
	ValueOfAnimal()
	ReflectCall()
	ReflectCallSlice()
}
