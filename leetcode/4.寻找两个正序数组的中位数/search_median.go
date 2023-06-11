package search_median

import (
	"fmt"
	"strconv"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := []int{}
	n1, n2 := 0, 0
	len1, len2 := len(nums1), len(nums2)
	for {
		if n1 == len1 || n2 == len2 || len1 == 0 || len2 == 0 {
			break
		}
		if nums1[n1] < nums2[n2] {
			merge = append(merge, nums1[n1])
			n1++
		} else {
			merge = append(merge, nums2[n2])
			n2++
		}
	}
	if n1 < len1 {
		merge = append(merge, nums1[n1:]...)
	}
	if n2 < len2 {
		merge = append(merge, nums2[n2:]...)
	}
	return median(merge)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func median(slice []int) float64 {
	length := len(slice)
	var res float64
	if length%2 == 0 {
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", float64(slice[length/2-1]+slice[length/2])/2), 64)
	} else {
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", float64(slice[length/2])), 64)
	}
	return res
}
