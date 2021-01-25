//package container_list
package main

//  双向链表
// go语言中list的实现原理是双向链表。list能高效的进行任意位置的元素插入和删除
//list是值类型，不不过采⽤用list的New()⽅方法声明的是⼀一个指针变量量。所以在
//拷⻉贝操作和参数传递时具有引⽤用类型的特征
import (
	"container/list"
	"fmt"
)

// 往链表头部插入一个元素，返回插入元素值的指针
//func (l *List) PushFront(v interface{}) *Element

func main() {

	// 创建一个链表，返回一个链表指针
	// func New() *List { return new(List).Init() }
	l := list.New()
	fmt.Println(l.Len())
	fmt.Println("------------------------")

	// 往链表中插入元素
	// 往链表头部插入一个元素，返回插入原的指针
	l.PushFront("go")
	// 往链表尾部插入一个元素，返回插入原的指针
	l.PushBack("java")
	l.PushBack("javaee")

	// 遍历链表
	// Front returns the first element of list l or nil if the list is empty.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("------------------------")

	// 返回链表的第一个元素
	front := l.Front()
	// 返回链表的最后一个元素
	back := l.Back()
	fmt.Println("first element: ", front.Value)
	fmt.Println("last element:  ", back.Value)
	fmt.Println("------------------------")

	// 往mark元素前插入一个元素
	l.InsertBefore("python", front)
	// 往mark元素后插入一个元素
	l.InsertAfter("vue", back)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("------------------------")

	//func (l *List) MoveToFront(e *Element) // 将元素移动到链表头部
	l.MoveToFront(back)
	//func (l *List) MoveToBack(e *Element) // 将元素移动到链表尾部
	l.MoveToBack(front)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("------------------------")

	l2 := list.New()
	a := l2.PushBack("a")
	b := l2.PushBack("b")
	c := l2.PushBack("c")
	d := l2.PushBack("d")

	// 顺序 a b c d
	//func (l *List) MoveBefore(e, mark *Element) // 将e元素移动到mark元素前面
	//func (l *List) MoveAfter(e, mark *Element) // 将e元素移动到mark元素后
	l2.MoveBefore(b, a) // b移动到a的前边
	l2.MoveAfter(c, d)  // c移动到的后边
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("------------------------")

	//func (l *List) Remove(e *Element) interface{} // 删除e元素
	l2.Remove(a)
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("------------------------")

	// 清空链表
	l2.Init()
	fmt.Println("清空l2链表", l2.Len())

	fmt.Println("------------------------")

	//func (l *List) PushFrontList(other *List) // 将other链表追加到l链表的头部
	//func (l *List) PushBackList(other *List) // 将other链表追加到l链表的末尾

	l3 := list.New()
	l3.PushBack("l3")

	l4 := list.New()
	l4.PushBack("l4")

	l5 := list.New()
	l5.PushBack("l5")

	l3.PushFrontList(l4)
	l5.PushBackList(l3)

	for e := l3.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("------------------------")

	for e := l4.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

	}

}
