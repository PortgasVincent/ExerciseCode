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

// func StackSort(stack *Stack) {
// 	subStack := &Stack{
// 		value: []int{},
// 	}
// 	for !stack.Empty() {
// 		first := stack.Top()
// 		stack.Pop()
// 		if stack.Empty() {
// 			subStack.Push(first)
// 			break
// 		}
// 		second := stack.Top()
// 		if first >= second {
// 			insertStack(first, stack, subStack)
// 		} else {
// 			stack.Pop()
// 			stack.Push(first)
// 			insertStack(second, stack, subStack)
// 		}
// 	}
// 	for !subStack.Empty() {
// 		stack.Push(subStack.Top())
// 		subStack.Pop()
// 	}
// 	return
// }

// func insertStack(val int, master, sub *Stack) {
// 	if !sub.Empty() {
// 		first := sub.Top()
// 		for val > first {
// 			master.Push(first)
// 			sub.Pop()
// 			if sub.Empty() {
// 				break
// 			}
// 			first = sub.Top()
// 		}
// 	}
// 	sub.Push(val)
// 	return
// }

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
		if stack.Empty() {
			stack.Push(first)
			break
		}
		second := stack.Top()
		if first >= second {
			insertStack(first, stack, subStack)
		} else {
			stack.Pop()
			stack.Push(first)
			insertStack(second, stack, subStack)
		}
	}
	for !subStack.Empty() {
		stack.Push(subStack.Top())
		subStack.Pop()
	}
}

func insertStack(val int, master, sub *Stack) {
	if sub.Empty() {
		sub.Push(val)
		return
	}
	first := sub.Top()
	for first < val {
		sub.Pop()
		master.Push(first)
		if sub.Empty() {
			break
		}
		first = sub.Top()
	}
	sub.Push(val)
	return
}
