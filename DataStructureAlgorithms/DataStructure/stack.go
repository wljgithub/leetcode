package main

import (
	"fmt"
	)

//  先进后出
func NewStack(size int) *Stack {
	if size <= 0 {
		panic("stack size must larger than 0")
	}
	return &Stack{
		Size: size}
}

type Stack struct {
	Size    int
	Content []interface{}
}

// 出栈
func (s *Stack) Pop() interface{} {
	v := s.Content[0]
	s.Content = s.Content[1:]
	fmt.Println("stack:",s.Content)
	return v
}

// 入栈
func (s *Stack) Push(value interface{}) {
	// 超出栈的长度,删掉最后几个元素，加上新的元素
	if len(s.Content) >= s.Size {
		fmt.Println("stack over len",s.Content)
		s.Content = s.Content[len(s.Content)-s.Size+1:]
	}
	s.Content = append(s.Content, value)

}

func (s *Stack) iterOver() {
	for _, v := range s.Content {
		fmt.Println(v)
	}
}

func main() {
	s := NewStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.iterOver()

	//fmt.Println(s.Pop())


}
