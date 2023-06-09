package two_sum

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标
*/

func twoSumEnumeration(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-x == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumHash(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, x := range nums {
		if result, ok := hashTable[target-x]; ok {
			return []int{result, i}
		}
		hashTable[x] = i
	}
	return nil
}
