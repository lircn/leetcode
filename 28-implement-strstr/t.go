package main

import "log"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	s1 := string("112233")
	s2 := string("23")
	log.Println(strStr(s1, s2))
}
