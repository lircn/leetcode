package main

import "log"
import "sort"

func lineEqual(a, b []int) bool {
	if a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] {
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

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	log.Println(nums)

	numMap := [][]int{}
	for i := 0; i <= len(nums)-4; i++ {
		for j := i + 1; j <= len(nums)-3; j++ {
			x := j + 1
			y := len(nums) - 1
			for x < y {
				if nums[i]+nums[j]+nums[x]+nums[y] == target {
					line := []int{nums[i], nums[j], nums[x], nums[y]}
					if !findLine(numMap, line) {
						numMap = append(numMap, line)
					}
					x++
					y--
				}
				if nums[i]+nums[j]+nums[x]+nums[y] > target {
					y--
				} else if nums[i]+nums[j]+nums[x]+nums[y] < target {
					x++
				}
			}
		}
	}
	return numMap
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	n := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	log.Println(fourSum(n, 3))
}
