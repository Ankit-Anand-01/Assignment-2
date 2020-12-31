//Stack implementation in Go
package main

import "fmt"

//Stack represents stack that hold a slice
type Stack struct {
	items []int
}

//Push
//Push will add the value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop
//Pop will remove the value at the end
//and returns it
func (s *Stack) Pop() int {
	le := len(s.items) - 1
	toRemove := s.items[le]
	s.items = s.items[:le]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(5)
	myStack.Push(15)
	myStack.Push(25)
	myStack.Push(35)
	myStack.Push(45)
	myStack.Push(55)
	myStack.Push(65)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
