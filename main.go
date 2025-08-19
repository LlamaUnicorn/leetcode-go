package main

import (
	"fmt"
)

type MyHashMap struct {
}

func Constructor() MyHashMap {

}

func (this *MyHashMap) Put(key int, value int) {

}

func (this *MyHashMap) Get(key int) int {

}

func (this *MyHashMap) Remove(key int) {

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

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
