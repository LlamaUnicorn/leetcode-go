package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

// LinkedList represents the linked list itself.
type LinkedList struct {
	Head *ListNode
}

// Add appends a new value to the end of the linked list.
func (list *LinkedList) Add(value int) {
	newNode := &ListNode{Value: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Display prints all the elements in the linked list.
func (list *LinkedList) Display() {
	current := list.Head
	elems := 0
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
		elems++
	}

	fmt.Println("nil")
	fmt.Println(elems)
}

func main() {
	list := &LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Display()
	//result := ll876.ListNode{}
	//fmt.Println(result)
}
