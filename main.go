package main

import (
	"fmt"
)

type MyHashSet struct {
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {

}

func (this *MyHashSet) Remove(key int) {

}

func (this *MyHashSet) Contains(key int) bool {
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main() {
	// Test cases
	//fmt.Println(reverseWords("the sky is blue"))  // blue is sky the
	//fmt.Println(reverseWords("  hello world  "))  // world hello
	//fmt.Println(reverseWords("a good   example")) // example good a
	fmt.Println("hey")
}
