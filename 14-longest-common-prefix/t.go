package main

import "log"

var INT_MIN int32 = int32(^INT_MAX)
var INT_MAX int = int(^uint(0) >> 1)

func longestCommonPrefix(strs []string) string {
	pre := ""
	stop := false
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		c := byte(' ')
		for _, s := range strs {
			if i >= len(s) {
				stop = true
				break
			}
			if c == ' ' {
				c = s[i]
			}
			if c != s[i] {
				stop = true
				break
			}
		}
		if stop {
			break
		} else {
			pre += string(c)
		}
	}
	return pre
}

func main() {
	log.SetFlags(log.Lshortfile)
	s := []string{"", ""}
	log.Println(s)
	log.Println(longestCommonPrefix(s))
}
