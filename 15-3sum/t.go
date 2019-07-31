package main

import "log"
import "sort"

func lineEqual(a, b []int) bool {
	if a[0] == b[0] && a[1] == b[1] && a[2] == b[2] {
		return true
	}
	return false
}

func findLine(numMap [][]int, line []int) bool {
	for _, v := range numMap {
		if lineEqual(v, line) {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	log.Println(nums)

	numMap := [][]int{}
	for k, v := range nums {
		x := k + 1
		y := len(nums) - 1
		for x < y {
			if v+nums[x]+nums[y] == 0 {
				line := []int{v, nums[x], nums[y]}
				if !findLine(numMap, line) {
					numMap = append(numMap, line)
				}
				x++
				y--
			}
			if v+nums[x]+nums[y] > 0 {
				y--
			} else if v+nums[x]+nums[y] < 0 {
				x++
			}
		}
	}

	return numMap
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	nn := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	log.Println(nn)
	log.Println("start")
	log.Println(threeSum(nn))
	log.Println("end")
}
