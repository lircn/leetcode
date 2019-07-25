package main

import "log"

var INT_MIN int32 = int32(^INT_MAX)
var INT_MAX int = int(^uint(0) >> 1)

func helper(n int, one, five, ten string) string {
	if n == 9 {
		return one + ten
	} else if n > 5 {
		s := five
		for n > 5 {
			s += one
			n--
		}
		return s
	} else if n == 5 {
		return five
	} else if n == 4 {
		return one + five
	} else if n > 0 {
		s := ""
		for n > 0 {
			s += one
			n--
		}
		return s
	} else {
		return ""
	}
}

func intToRoman(num int) string {
	s := ""
	if num >= 1000 {
		s += helper(num/1000, "M", "?", "!")
		num %= 1000
	}
	if num >= 100 {
		s += helper(num/100, "C", "D", "M")
		num %= 100
	}
	if num >= 10 {
		s += helper(num/10, "X", "L", "C")
		num %= 10
	}
	s += helper(num, "I", "V", "X")
	return s
}

func main() {
	log.SetFlags(log.Lshortfile)
	n := 58
	log.Println(n)
	log.Println(intToRoman(n))
}
