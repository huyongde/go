package structure

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Print()
	s.Top()
	s.Pop()
	s.Push(1)
	s.Push(2)
	s.Print()
	e := s.Top()
	fmt.Println(e)
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
}
