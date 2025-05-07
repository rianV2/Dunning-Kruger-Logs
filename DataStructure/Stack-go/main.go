package main

import "fmt"

type Stack struct {
	items []int
}

// push to end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// remove end and return end
func (s *Stack) Pop() int {
	indEnd := len(s.items) - 1
	toRemove := s.items[indEnd]
	s.items = s.items[0:indEnd]

	return toRemove
}

// return end
func (s *Stack) Peek() int {
	indEnd := len(s.items) - 1

	return s.items[indEnd]
}

// check empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	slices := Stack{}
	slices.Push(1)
	slices.Push(2)
	slices.Push(3)
	fmt.Println(slices)

	fmt.Println("------")
	slices.Pop()
	slices.Pop()
	slices.Peek()
	fmt.Println(slices)
	fmt.Println("------")
	slices.Pop()
	slices.IsEmpty()
	fmt.Println(slices)

}
