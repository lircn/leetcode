package main

import "log"
import "sort"

func delta(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	log.Println(nums)

	minDelta := int(^uint(0) >> 1)
	var ret int
	for k, v := range nums {
		x := k + 1
		y := len(nums) - 1
		for x < y {
			sum := v + nums[x] + nums[y]
			d := delta(sum, target)
			if minDelta > d {
				minDelta = d
				ret = sum
			}
			if sum > target {
				y--
			} else if sum < target {
				x++
			} else if sum == target {
				return sum
			}
		}
	}
	return ret
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	nn := []int{0, 2, 1, -3}
	log.Println("start")
	log.Println(threeSumClosest(nn, 1))
	log.Println("end")
}
