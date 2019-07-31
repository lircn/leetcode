package main

import "log"
import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := list.New()
	ori := head
	length := 0
	for head != nil {
		length++
		l.PushBack(head)
		if l.Len() > n+1 {
			l.Remove(l.Front())
		}
		head = head.Next
	}
	rm := l.Front().Value.(*ListNode)
	if length == n {
		return rm.Next
	} else {
		rm.Next = rm.Next.Next
		return ori
	}
}

func addNode(node *ListNode, n int) *ListNode {
	node.Next = &ListNode{n, nil}
	return node.Next
}

func initList() *ListNode {
	head := &ListNode{1, nil}
	node := head
	node = addNode(node, 2)
	node = addNode(node, 3)
	return head
}

func printList(head *ListNode) {
	for head != nil {
		log.Print(head.Val)
		head = head.Next
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	head := initList()
	printList(head)
	head = removeNthFromEnd(head, 1)
	printList(head)
}
