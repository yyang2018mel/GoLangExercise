package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	next := head.Next

	for next != nil {
		if current.Val == next.Val {
			current.Next = next.Next
			next = next.Next
		} else {
			current = next
			next = next.Next
		}
	}

	return head
}

func makeListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := new(ListNode)
	head.Val = nums[0]

	current := head
	for _, val := range nums[1:] {
		newNode := new(ListNode)
		newNode.Val = val
		current.Next = newNode
		current = current.Next
	}

	return head
}

func main() {
	arr := []int{1, 2, 3, 3}
	list := makeListNode(arr)
	reducedList := deleteDuplicates(list)
	fmt.Println(reducedList)
}
