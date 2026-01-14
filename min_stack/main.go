package main

import "fmt"

type MinStack struct {
	topValue []int
	minValue []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	if len(s.topValue) == 0 {
		s.topValue = append(s.topValue, val)
		s.minValue = append(s.minValue, val)
		return
	}
	m := s.minValue[len(s.topValue)-1]
	if m > val {
		s.minValue = append(s.minValue, val)
		s.topValue = append(s.topValue, val)
		return
	}
	s.minValue = append(s.minValue, m)
	s.topValue = append(s.topValue, val)
}

func (s *MinStack) Pop() {
	s.topValue = s.topValue[:len(s.topValue)-1]
	s.minValue = s.minValue[:len(s.minValue)-1]
}

func (s *MinStack) Top() int {
	return s.topValue[len(s.topValue)-1]
}

func (s *MinStack) GetMin() int {
	return s.minValue[len(s.minValue)-1]
}

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	fmt.Println(s.GetMin())
	s.Push(-3)
	s.Push(2)
	fmt.Println(s.GetMin())
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
}
