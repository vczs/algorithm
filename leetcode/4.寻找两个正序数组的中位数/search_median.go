package search_median

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	merge := []int{}

	n1, n2 := 0, 0
	for {
		if n1 == len1 || n2 == len2 {
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

func median(slice []int) float64 {
	length := len(slice)
	if length%2 == 0 {
		return float64(slice[length/2-1]+slice[length/2]) / 2
	} else {
		return float64(slice[length/2])
	}
}
