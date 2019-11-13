package main

import "log"

var INT_MAX int32 = int32(^uint32(0) >> 1)
var INT_MIN int32 = int32(^INT_MAX)

func divide(dividend int32, divisor int32) int32 {
	flag := int32(1)
	if dividend == divisor {
		return 1
	}
	if dividend < 0 {
		flag = -flag
		dividend = -dividend
	}
	if divisor < 0 {
		flag = -flag
		divisor = -divisor
	}
	n := int32(0)
	for dividend >= divisor {
		dividend -= divisor
		n++
	}
	return n * flag
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	log.Println(divide(-2147483648, -1))
	log.Println(INT_MAX)
	log.Println(INT_MIN)
}
