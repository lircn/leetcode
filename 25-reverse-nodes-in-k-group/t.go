package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	root := &ListNode{0, head}
	prev := root
	curr := head
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	for i := 0; i < length/k; i++ {
		for j := 0; j < k-1; j++ {
			next := curr.Next
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		prev = curr
		curr = prev.Next
	}
	return root.Next
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
	l := initList([]int{1, 2, 3, 4, 5})
	printList("before", l)
	printList("after", reverseKGroup(l, 3))
}
