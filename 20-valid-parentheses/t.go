package main

import "log"
import "container/list"

func isValidHelper(l, r rune) bool {
	if l == '(' {
		return r == ')'
	}
	if l == '[' {
		return r == ']'
	}
	if l == '{' {
		return r == '}'
	}
	return false
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	l := list.New()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			l.PushBack(c)
		} else {
			if l.Len() == 0 {
				return false
			}
			back := l.Back()
			if !isValidHelper(back.Value.(rune), c) {
				return false
			}
			l.Remove(back)
		}
	}
	if l.Len() == 0 {
		return true
	} else {
		return false
	}
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		log.Println(e.Value.(rune))
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	s := ")]}"
	log.Println(s)
	log.Println(isValid(s))
}
