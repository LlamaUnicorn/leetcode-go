package main

import (
	"fmt"
	"github.com/LlamaUnicorn/leetcode-go/linked-list/ll876"
)

type SingleListNode struct {
	Val  int
	Next *SingleListNode
}
func main() {

	FirstNode := SingleListNode{
		Val: 1,
		Next: &SingleListNode{SecondNode}
	}
	SecondNode := SingleListNode{
		val: 2,
		Next: &SingleListNode{}
	}
	result := ll876.ListNode{}
	fmt.Println(result)
}
