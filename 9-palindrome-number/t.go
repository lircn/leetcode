package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	ori := x
	n := 0
	for x > 0 {
		n += x % 10
		x /= 10
		n *= 10
	}
	n /= 10
	if n == ori {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("vim-go")
}
