package main

import "fmt"

func reverse(x int) int {
	const INT_MAX = int32(^uint32(0) >> 1)
	const INT_MIN = int32(^INT_MAX)
	ret := int(0)
	for {
		ret += x % 10
		x /= 10
		if x == 0 {
			break
		}
		ret *= 10
	}
	if ret > int(INT_MAX) {
		return 0
	}
	if ret < int(INT_MIN) {
		return 0
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}
