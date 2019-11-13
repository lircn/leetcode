package main

import "log"

func removeElement(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		} else {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	n := []int{1, 1, 2, 2, 3, 4, 4}
	log.Println(n)
	log.Println(removeElement(n, 2))
	log.Println(n)
}
