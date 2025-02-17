package ll876

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Value int
	Next  *ListNode
}

// Function to find the middle node of the linked list
func middleNode(head *ListNode) *ListNode {
	// Initialize slow and fast pointers
	slow, fast := head, head

	// Traverse the list using the two-pointer technique
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move slow pointer one step
		fast = fast.Next.Next // Move fast pointer two steps
	}

	// When fast reaches the end, slow will be at the middle
	return slow
}

// Helper function to create a linked list from a slice of integers
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Value: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Value: values[i]}
		current = current.Next
	}
	return head
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Example 1: [1, 2, 3, 4, 5]
	head1 := createLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Print("Middle of [1, 2, 3, 4, 5]: ")
	printLinkedList(middleNode(head1)) // Output: [3, 4, 5]

	// Example 2: [1, 2, 3, 4, 5, 6]
	head2 := createLinkedList([]int{1, 2, 3, 4, 5, 6})
	fmt.Print("Middle of [1, 2, 3, 4, 5, 6]: ")
	printLinkedList(middleNode(head2)) // Output: [4, 5, 6]
}
