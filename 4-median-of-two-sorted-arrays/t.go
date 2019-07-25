package main

import "fmt"

func cut2n2(nums1 []int, nums2 []int) ([]int, []int) {
	if nums1[0] < nums2[0] {
		nums1 = nums1[1:]
	} else {
		nums2 = nums2[1:]
	}
	if nums1[len(nums1)-1] < nums2[len(nums2)-1] {
		nums2 = nums2[:len(nums2)-1]
	} else {
		nums1 = nums1[:len(nums1)-1]
	}
	return nums1, nums2
}

/* len(nums1) == 1 */
func cut1n2(nums1 []int, nums2 []int) ([]int, []int) {
	if nums1[0] < nums2[0] {
		nums1 = make([]int, 0)
		nums2 = nums2[:len(nums2)-1]
	} else {
		nums2 = nums2[1:]
		if nums1[len(nums1)-1] < nums2[len(nums2)-1] {
			nums2 = nums2[:len(nums2)-1]
		} else {
			nums1 = make([]int, 0)
		}
	}
	return nums1, nums2
}

/* len(nums1) == 0 */
func cut0n2(nums1 []int, nums2 []int) ([]int, []int) {
	nums2 = nums2[1 : len(nums2)-1]
	return nums1, nums2
}

func getRet(nums1 []int, nums2 []int) float64 {
	sum := float64(0)
	cnt := float64(0)
	for _, v := range nums1 {
		sum += float64(v)
		cnt += 1
	}
	for _, v := range nums2 {
		sum += float64(v)
		cnt += 1
	}
	if cnt > 0 {
		return sum / cnt
	} else {
		return 0
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for {
		if len(nums1)+len(nums2) <= 2 {
			return getRet(nums1, nums2)
		}

		if len(nums1) == 0 {
			nums1, nums2 = cut0n2(nums1, nums2)
			continue
		}
		if len(nums2) == 0 {
			nums1, nums2 = cut0n2(nums2, nums1)
			continue
		}

		if len(nums1) == 1 {
			nums1, nums2 = cut1n2(nums1, nums2)
			continue
		}
		if len(nums2) == 1 {
			nums1, nums2 = cut1n2(nums2, nums1)
			continue
		}

		if len(nums1) >= 2 && len(nums2) >= 2 {
			nums1, nums2 = cut2n2(nums1, nums2)
			continue
		}

	}
	return -1
}

func main() {
	fmt.Println("vim-go")
}
