//Linked_list in Golang
package main

import "fmt"

type list struct {
	digi     int
	prevList *list
}

// push adds another number to the stack
func (l *list) push(num int) *list {
	var current list
	current.digi = num
	current.prevList = l
	return &current
}

// pop removes the top most member from the stack
func (l *list) pop() {
	l.digi = l.prevList.digi
	l.prevList = l.prevList.prevList
}

func main() {
	var top *list
	// Create parts 0-9 of the stack
	for i := 0; i < 10; i++ {
		top = top.push(i)
	}

	top = top.push(10)
	fmt.Println(*top.prevList.prevList)
	fmt.Println(*top.prevList)
	fmt.Println(*top)
	top.pop()
	fmt.Println(*top)
}
