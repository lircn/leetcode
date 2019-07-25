package main

import "fmt"

var INT_MAX int32 = int32(^uint32(0) >> 1)
var INT_MIN int32 = int32(^INT_MAX)

func toInt(str string) int {
	ret := int(0)
	for _, c := range str {
		if c >= '0' && c <= '9' {
			ret += int(c) - int('0')
		} else {
			break
		}
		ret *= 10
		if ret > int(INT_MAX)*100 {
			break
		}
	}
	ret /= 10
	return ret
}

func myAtoiHelper(str string) (int, int) {
	for i, c := range str {
		if c == ' ' {
			continue
		}
		str = str[i:]
		break
	}

	for i, c := range str {
		if c == '-' {
			return -1, -1 * toInt(str[i+1:])
		}
		if c == '+' {
			return 1, toInt(str[i+1:])
		}
		if c >= '0' && c <= '9' {
			return 1, toInt(str[i:])
		}
		return 1, 0
	}
	return 1, 0
}

func myAtoi(str string) int {
	flag, ret := myAtoiHelper(str)

	if ret < int(INT_MIN) || ret > int(INT_MAX) {
		if flag < 0 {
			return int(INT_MIN)
		} else {
			return int(INT_MAX)
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}
