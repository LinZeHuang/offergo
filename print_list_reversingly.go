package just_to_offer

import (
	"errors"
	"fmt"
)

//这是单向链表结构
type MyList struct {
	Value interface{}
	Next  *MyList
}

//将链表反过来之后去打印，修改了链表结构，仍然不是纯函数（其实我非常讨厌非纯函数）
func PrintListReversingly(list *MyList) *MyList {
	if list == nil {
		return nil
	}
	var par *MyList
	for list != nil {
		next := list.Next
		list.Next = par
		par = list
		list = next
	}
	return par
}

//利用栈去输出链表。
type Stack struct {
	Elem []interface{}
}

func NewStack() *Stack {
	return &Stack{
		Elem: make([]interface{}, 0),
	}
}

func (self *Stack) Pop() interface{} {
	if self == nil || len(self.Elem) <= 0 {
		return nil
	}
	elem := self.Elem[len(self.Elem)-1]
	self.Elem = self.Elem[:len(self.Elem)-1]
	return elem
}

func (self *Stack) Push(elem interface{}) error {
	if self == nil {
		return errors.New("nil stack")
	}
	self.Elem = append(self.Elem, elem)
	return nil
}

func PrintListReverByStack(list *MyList) {
	if list == nil {
		return
	}
	stack := NewStack()
	//避免修改到原list
	myList := list
	if myList != nil {
		stack.Push(list)
		myList = myList.Next
	}
	for {
		elem := stack.Pop()
		if elem == nil {
			break
		}
		fmt.Printf("%v\n", elem)
	}

}

//利用递归
func PrintListReverByRecursion(list *MyList) {
	if list != nil {
		if list.Next != nil {
			PrintListReverByRecursion(list.Next)
		}
		fmt.Printf("%v\n", list.Value)
	}
}
