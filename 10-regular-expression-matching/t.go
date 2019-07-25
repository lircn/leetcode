package main

import "log"

func helper(s, p string) (string, string) {
	star := p[len(p)-2]
	if star == '.' {
		star = s[len(s)-1]
	}

	if s[len(s)-1] != star {
		return s, p[:len(p)-2]
	}

	for len(s) > 0 {
		if s[len(s)-1] == star {
			s = s[:len(s)-1]
		} else {
			break
		}
	}

	p = p[:len(p)-2]
	for len(p) > 0 {
		if p[len(p)-1] == star {
			p = p[:len(p)-1]
		} else {
			break
		}
	}

	return s, p
}

func checkP(p string) string {
	for len(p) >= 2 {
		if p[len(p)-1] == '*' {
			p = p[:len(p)-2]
		} else {
			break
		}
	}
	return p
}

func isMatch(s string, p string) bool {
	for {
		if len(p) <= 1 || len(s) == 0 {
			break
		}

		pLast := len(p) - 1
		sLast := len(s) - 1
		if p[pLast] == '.' || p[pLast] == s[sLast] {
			s = s[:sLast]
			p = p[:pLast]
			continue
		} else if p[pLast] == '*' {
			if pLast == 0 {
				p = p[:pLast]
			} else {
				s, p = helper(s, p)
			}
			continue
		}

		log.Println(s, ", ", p)
		if p[len(p)-1] != s[len(s)-1] {
			return false
		}
	}

	if len(p) > 1 {
		p = checkP(p)
	}

	log.Println(s, ", ", p)
	if len(p) == len(s) {
		return true
	} else {
		return false
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	s := "ab"
	p := ".*"
	log.Println(s, ", ", p)
	log.Println(isMatch(s, p))
}
