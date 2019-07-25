package main

import "fmt"

func getSubStr(s string, a, b int) string {
	if a+1 < b {
		return s[a+1 : b]
	} else {
		return ""
	}
}

func backFind(s string, pos int) string {
	a := pos
	b := pos + 1
	for a >= 0 && b < len(s) {
		if s[a] != s[b] {
			break
		}
		a--
		b++
	}
	s1 := getSubStr(s, a, b)

	a = pos - 1
	b = pos + 1
	for a >= 0 && b < len(s) {
		if s[a] != s[b] {
			break
		}
		a--
		b++
	}
	s2 := getSubStr(s, a, b)

	if len(s1) > len(s2) {
		return s1
	} else {
		return s2
	}
}

func frontFind(s string, pos int) string {
	a := pos - 1
	b := pos
	for a >= 0 && b < len(s) {
		if s[a] != s[b] {
			break
		}
		a--
		b++
	}
	s1 := getSubStr(s, a, b)

	a = pos - 1
	b = pos + 1
	for a >= 0 && b < len(s) {
		if s[a] != s[b] {
			break
		}
		a--
		b++
	}
	s2 := getSubStr(s, a, b)

	if len(s1) > len(s2) {
		return s1
	} else {
		return s2
	}
}

func longestPalindrome(s string) string {
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[1:]
		}
	}
	if len(s) <= 1 {
		return s
	}

	max := ""
	mid := len(s) / 2
	for mid > 0 {
		cnt := frontFind(s, mid)
		if len(cnt) > len(max) {
			max = cnt
		}
		if len(cnt) <= 0 {
			break
		}
		mid--
	}

	mid = len(s) / 2
	for mid < len(s)-1 {
		cnt := backFind(s, mid)
		if len(cnt) > len(max) {
			max = cnt
		}
		if len(cnt) <= 0 {
			break
		}
		mid++
	}

	return max
}

func main() {
	fmt.Println("vim-go")
}
