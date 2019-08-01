package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPair(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPair(next.Next)
	next.Next = head
	return next
}

func swapPairs(head *ListNode) *ListNode {
	return swapPair(head)
}

func addNode(head *ListNode, v int) *ListNode {
	if head == nil {
		return &ListNode{v, nil}
	}
	ori := head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{v, nil}
	return ori
}

func initList(l []int) *ListNode {
	var head *ListNode
	for _, i := range l {
		head = addNode(head, i)
	}
	return head
}

func printList(tag string, head *ListNode) {
	for i := 0; head != nil; i++ {
		log.Print(tag, ",", i, ": ", head.Val)
		head = head.Next
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	l := initList([]int{1, 2, 3, 4})
	printList("before", l)
	printList("after", swapPairs(l))
}
