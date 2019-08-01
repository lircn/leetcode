package main

import "log"

func gen(l, r int, s string, tab *[]string) {
	if l == 0 && r == 0 {
		*tab = append(*tab, s)
		return
	}
	if l == 0 {
		tmp := ""
		for ; r > 0; r-- {
			tmp += ")"
		}
		gen(0, 0, s+tmp, tab)
		return
	}
	if l < r {
		gen(l, r-1, s+")", tab)
		gen(l-1, r, s+"(", tab)
	} else {
		gen(l-1, r, s+"(", tab)
	}
}

func generateParenthesis(n int) []string {
	var tab []string
	gen(n, n, "", &tab)
	return tab
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	log.Println(generateParenthesis(4))
}
