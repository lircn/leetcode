package main

import "fmt"

func convert2(s string) string {
	s1 := ""
	s2 := ""
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			s1 += string(s[i])
		} else {
			s2 += string(s[i])
		}
	}
	return s1 + s2
}
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	if numRows == 2 {
		return convert2(s)
	}

	data := make([]string, numRows)
	down := true
	pos := 0
	for i := 0; i < len(s); i++ {
		if down {
			if pos >= numRows {
				pos -= 2
				down = false
			}
		} else {
			if pos == 0 {
				down = true
			}
		}
		data[pos] += string(s[i])

		if down {
			pos++
		} else {
			pos--
		}
	}

	retStr := ""
	for i := 0; i < numRows; i++ {
		retStr += data[i]
	}
	return retStr
}

func main() {
	fmt.Println("vim-go")
}
