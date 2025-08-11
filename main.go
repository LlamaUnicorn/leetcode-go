package main

import (
	"fmt"
)

type MyHashSet struct {
	present []bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		present: make([]bool, 1_000_001),
	}
}

func (this *MyHashSet) Add(key int) {
	this.present[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.present[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.present[key]
}

func main() {
	// Test cases
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	fmt.Println(obj.Contains(1))
	fmt.Println(obj.Contains(3))
	obj.Add(2)
	fmt.Println(obj.Contains(2))
	obj.Remove(2)
	obj.Contains(2)
	fmt.Println(obj.Contains(2))
}
