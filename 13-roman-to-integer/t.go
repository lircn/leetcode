package main

import "log"

func romanToInt(s string) int {
	sum := 0
	for len(s) >= 2 {
		if s[len(s)-1] == 'M' {
			if s[len(s)-2] == 'C' {
				sum += 900
				s = s[:len(s)-2]
			} else {
				sum += 1000
				s = s[:len(s)-1]
			}
			continue
		}
		if s[len(s)-1] == 'D' {
			if s[len(s)-2] == 'C' {
				sum += 400
				s = s[:len(s)-2]
			} else {
				sum += 500
				s = s[:len(s)-1]
			}
			continue
		}
		if s[len(s)-1] == 'C' {
			if s[len(s)-2] == 'X' {
				sum += 90
				s = s[:len(s)-2]
			} else {
				sum += 100
				s = s[:len(s)-1]
			}
			continue
		}
		if s[len(s)-1] == 'L' {
			if s[len(s)-2] == 'X' {
				sum += 40
				s = s[:len(s)-2]
			} else {
				sum += 50
				s = s[:len(s)-1]
			}
			continue
		}
		if s[len(s)-1] == 'X' {
			if s[len(s)-2] == 'I' {
				sum += 9
				s = s[:len(s)-2]
			} else {
				sum += 10
				s = s[:len(s)-1]
			}
			continue
		}
		if s[len(s)-1] == 'V' {
			if s[len(s)-2] == 'I' {
				sum += 4
				s = s[:len(s)-2]
			} else {
				sum += 5
				s = s[:len(s)-1]
			}
			continue
		}
		if s[len(s)-1] == 'I' {
			sum += 1
			s = s[:len(s)-1]
			continue
		}
	}

	if len(s) > 0 {
		switch s[0] {
		case 'I':
			sum += 1
		case 'V':
			sum += 5
		case 'X':
			sum += 10
		case 'L':
			sum += 50
		case 'C':
			sum += 100
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
	}
	return sum
}

func main() {
	log.SetFlags(log.Lshortfile)
	s := "MCMXCIV"
	log.Println(romanToInt(s))
}
