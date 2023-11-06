package main

import "fmt"

type Stack struct {
	value []int
}

func (s *Stack) Push(val int) {
	s.value = append(s.value, val)
	return
}

func Push[T any](val T, arr []T) {
	arr = append(arr, val)
}

func (s *Stack) Top() int {
	if len(s.value) == 0 {
		return 0
	}
	return s.value[len(s.value)-1]
}
func (s *Stack) Pop() {
	if len(s.value) == 0 {
		return
	}
	s.value = s.value[:len(s.value)-1]
	return
}
func (s *Stack) Empty() bool {
	return len(s.value) == 0
}

func main() {
	a := &Stack{
		value: []int{2, 3, 5, 12, 3, 2, 65, 7, 4, 1, 4, 2, 5},
	}
	StackSort(a)
	fmt.Println(a.value)
}

func StackSort(stack *Stack) {
	subStack := &Stack{
		value: []int{},
	}
	for !stack.Empty() {
		first := stack.Top()
		stack.Pop()

		insertStack(first, stack, subStack)
	}
	for !subStack.Empty() {
		stack.Push(subStack.Top())
		subStack.Pop()
	}
	return
}

func insertStack(val int, master, sub *Stack) {
	if sub.Empty() {
		sub.Push(val)
		return
	}
	for val > sub.Top() {
		master.Push(sub.Top())
		sub.Pop()
		if sub.Empty() {
			break
		}
	}
	sub.Push(val)
	return
}
