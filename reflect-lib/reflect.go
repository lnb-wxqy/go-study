package reflect_lib

import (
	"fmt"
	"reflect"
)

// 获取对象的类型信息
//func reflect.TypeOf(i interface{}) reflect.Type

type A int

// ---------------使用反射获取对象的类型-----------------

func TypeOfMethod() {
	var a A = 1
	t := reflect.TypeOf(a)
	fmt.Println(t.Name(), t.Kind()) //A int； Name 是真实类型，Kind是真是类型的基础结构类型
}

//传入的参数对象区分基类型和指针类型
func TypeOfMethod2() {
	b := 2
	t1, t2 := reflect.TypeOf(b), reflect.TypeOf(&b)
	// Elem方法返回指针的基类型，此外，Elem方法还可以返回数组、slice、map、chan的基类型
	fmt.Println(t1, t2, t1 == t2, t2.Elem()) //int *int false int;
	fmt.Println(t1.Kind(), t2.Kind())        // int ptr
	fmt.Println(t1 == t2.Elem())             //true
}

// --------------------使用反射操作结构体-----------------------
// 遍历结构体： 遍历结构体字段，需要先获取结构体指针的类型

type Student struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func ReflectStruct() {
	var s Student
	t := reflect.TypeOf(&s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fmt.Println(t.NumField()) // 2
	// 遍历结构体属性
	for i := 0; i < t.NumField(); i++ {
		sField := t.Field(i)
		// 获取tag
		// 输出结果：
		// name string name
		// age int age
		fmt.Println(sField.Name, sField.Type, sField.Tag.Get("json"))
	}
}

//使用反射获取和修改对象的值
//func reflect.ValueOf(i interface{}) reflect.Value
type animal struct {
	Name string
	age  int
}

func ValueOfAnimal() {
	var a animal
	//我们知道修改对象的变量，需要传入指针，但是因为接口变量存储的指针是不可以寻址和设置变量的，所以还需要通过 Elem 获取目标对象。
	v := reflect.ValueOf(&a).Elem()
	fmt.Printf("v: %v type: %v kind: %v\n", v, v.Type(), v.Kind()) // v: { 0} type: reflect_lib.animal kind: struct
	v.FieldByName("Name").SetString("kity")
	fmt.Printf("v: %v type: %v kind: %v\n", v, v.Type(), v.Kind()) //v: {kity 0} type: reflect_lib.animal kind: struct

}

// -------------------使用反射动态调用方法-----------------------
//固定参数，使用 Call，只需按照参数列表传参即可。
type member struct {
}

func (m member) MemberInfo(name string, score int) string {
	mInfo := fmt.Sprintf("姓名：%s, 积分：%d\n", name, score)
	return mInfo
}

func ReflectCall() {
	var m1 member
	v := reflect.ValueOf(&m1)
	m := v.MethodByName("MemberInfo")
	param := []reflect.Value{
		reflect.ValueOf("五行缺雨"),
		reflect.ValueOf(88),
	}

	result := m.Call(param)
	for _, r := range result {
		fmt.Println(r)
	}
}

// 可变参数，使用CallSlice，仅需要一个[]interface{}即可
type person struct {
}

func (p person) Hobby(name string, hobby ...interface{}) string {
	myHobby := fmt.Sprintf("I am %s,my hobby is: %v", name, hobby)
	return myHobby
}

func ReflectCallSlice() {
	var p1 person
	v := reflect.ValueOf(&p1)
	m := v.MethodByName("Hobby") // 姓名：五行缺雨, 积分：88

	param := []reflect.Value{
		reflect.ValueOf("五行缺雨"),
		reflect.ValueOf("dance"),
		reflect.ValueOf("sleep"),
	}
	result := m.Call(param)
	fmt.Println(result) // 	[I am 五行缺雨,my hobby is: [dance sleep]]

	param = []reflect.Value{
		reflect.ValueOf("五行缺雨"),
		reflect.ValueOf([]interface{}{
			"sleep", "singing",
		}),
	}
	result = m.CallSlice(param)
	fmt.Println(result) // [I am 五行缺雨,my hobby is: [sleep singing]]
}
