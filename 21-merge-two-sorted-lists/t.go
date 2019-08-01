package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	for !(l1 == nil && l2 == nil) {
		if l1 == nil {
			n := l2.Next
			l2.Next = nil
			l3 = addNode(l3, l2)
			l2 = n
			continue
		}
		if l2 == nil {
			n := l1.Next
			l1.Next = nil
			l3 = addNode(l3, l1)
			l1 = n
			continue
		}
		if l1.Val < l2.Val {
			n := l1.Next
			l1.Next = nil
			l3 = addNode(l3, l1)
			l1 = n
			continue
		} else {
			n := l2.Next
			l2.Next = nil
			l3 = addNode(l3, l2)
			l2 = n
			continue
		}
	}
	return l3
}

func addNode(head, node *ListNode) *ListNode {
	if head == nil {
		return node
	}
	ori := head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
	return ori
}
func addNodeV(node *ListNode, n int) *ListNode {
	node.Next = &ListNode{n, nil}
	return node.Next
}

func initList() *ListNode {
	head := &ListNode{1, nil}
	node := head
	node = addNodeV(node, 2)
	node = addNodeV(node, 3)
	return head
}

func printList(tag string, head *ListNode) {
	for head != nil {
		log.Print(tag, ": ", head.Val)
		head = head.Next
	}
}
func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	l1 := initList()
	l2 := initList()
	l3 := mergeTwoLists(l1, l2)
	printList("l3", l3)
}
