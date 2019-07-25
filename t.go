package main

import "log"

var INT_MIN int32 = int32(^INT_MAX)
var INT_MAX int = int(^uint(0) >> 1)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("123")
}
