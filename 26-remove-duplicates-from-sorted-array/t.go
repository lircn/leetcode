package main

import "log"

func removeDuplicates(nums []int) int {
	n := 0
	for i := 1; i < len(nums); i++ {
		if nums[n] == nums[i] {
			continue
		} else {
			n++
			nums[n] = nums[i]
		}
	}
	return n + 1
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	n := []int{1, 1, 2, 2, 3, 4, 4}
	log.Println(n)
	log.Println(removeDuplicates(n))
	log.Println(n)
}
