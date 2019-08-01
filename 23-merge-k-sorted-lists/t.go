package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

var INT_MAX int = int(^uint(0) >> 1)

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

func mergeKLists(lists []*ListNode) *ListNode {
	var sum *ListNode
	for len(lists) > 0 {
		min := INT_MAX
		var k int
		nilK := -1
		for lk, lv := range lists {
			if lv == nil {
				nilK = lk
				continue
			}
			if min > lv.Val {
				min = lv.Val
				k = lk
			}
		}
		if min != INT_MAX {
			sum = addNode(sum, min)
			lists[k] = lists[k].Next
		}
		if nilK != -1 {
			lists = append(lists[:nilK], lists[nilK+1:]...)
		}
	}
	return sum
}

func initList(l []int) *ListNode {
	var head *ListNode
	for _, i := range l {
		head = addNode(head, i)
	}
	return head
}

func initLists() []*ListNode {
	var l []*ListNode
	l = append(l, initList([]int{}))
	l = append(l, initList([]int{}))
	l = append(l, initList([]int{}))
	return l
}

func printList(tag string, head *ListNode) {
	for i := 0; head != nil; i++ {
		log.Print(tag, ",", i, ": ", head.Val)
		head = head.Next
	}
}

func printLists(tag string, lists []*ListNode) {
	for _, l := range lists {
		printList(tag, l)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	l := initLists()
	printList("m", mergeKLists(l))
}
