package main

import "log"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			n := min(height[i], height[j]) * (j - i)
			if n > max {
				max = n
			}
		}
	}
	return max
}

func main() {
	log.SetFlags(log.Lshortfile)
}
