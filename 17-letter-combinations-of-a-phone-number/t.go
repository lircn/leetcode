package main

import "log"
import "bytes"
import "container/list"

type Node struct {
	c      byte
	parent *Node
	child  []*Node
}

var KMap map[byte][]byte = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func addLeafs(node *Node, leafs []byte) {
	for _, leaf := range leafs {
		node.child = append(node.child, &Node{leaf, node, nil})
	}
}

func addNumToTree(root *Node, num byte) {
	if root.child == nil {
		addLeafs(root, KMap[num])
		return
	}

	for _, node := range root.child {
		addNumToTree(node, num)
	}
}

func leafToString(leaf *Node) string {
	l := list.New()
	for leaf.parent != nil {
		l.PushFront(leaf.c)
		leaf = leaf.parent
	}
	var b bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		b.WriteByte(e.Value.(byte))
	}
	return b.String()
}

func treeToStrings(root *Node, strs *[]string) {
	if root.child == nil {
		*strs = append(*strs, leafToString(root))
		return
	}
	for _, node := range root.child {
		treeToStrings(node, strs)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	root := Node{0, nil, nil}
	dBytes := []byte(digits)

	for _, c := range dBytes {
		addNumToTree(&root, c)
	}
	var strs []string
	treeToStrings(&root, &strs)
	return strs
}

func main() {
	log.SetFlags(log.Lshortfile)
	s := "23"
	log.Println("str: ", s)
	log.Println(letterCombinations(s))
}
