package structure

import (
	"fmt"
	"sync"
)

// 实现协程安全的栈
type stack struct {
	elem []interface{}
	sync.Mutex
}

func NewStack() *stack {
	s := new(stack)
	s.Mutex = sync.Mutex{}
	return s
}

func (s *stack) Push(v ...interface{}) {
	s.Lock()
	defer s.Unlock()
	s.elem = append(s.elem, v...)
}

// 返回栈顶元素
func (s *stack) Top() interface{} {
	s.Lock()
	defer s.Unlock()
	if s.Size() == 0 {
		return nil
	}
	return s.elem[s.Size()-1]
}

// 删除栈顶元素
func (s *stack) Pop() {
	s.Lock()
	defer s.Unlock()
	if s.Size() == 0 {
		return
	}
	s.elem = s.elem[0 : s.Size()-1]
}

// 判断栈是否为空
func (s *stack) Empty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}

//返回栈内元素个数
func (s *stack) Size() int {
	return len(s.elem)
}
func (s *stack) Print() {
	fmt.Println("start -----------")
	defer fmt.Println("end -----------")
	if s.Size() == 0 {
		return
	}
	for i := s.Size() - 1; i >= 0; i-- {
		fmt.Println(i, s.elem[i])
	}
}
