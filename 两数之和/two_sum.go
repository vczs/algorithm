package two_sum

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
*/

/**
暴力枚举:
最容易想到的方法是枚举数组中的每一个数 x，寻找数组中是否存在 target - x。
当我们使用遍历整个数组的方式寻找 target - x 时，需要注意到每一个位于 x 之前的元素都已经和 x 匹配过，因此不需要再进行匹配。
而每一个元素不能被使用两次，所以我们只需要在 x 后面的元素中寻找 target - x。

时间复杂度：O(N^2),其中 N 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
空间复杂度：O(1)。
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

/**
哈希表:
注意到方法一的时间复杂度较高的原因是寻找 target - x 的时间复杂度过高。
因此，我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。
使用哈希表，可以将寻找 target - x 的时间复杂度降低到从O(N)降低到O(1)。
这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，
然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。

时间复杂度：O(N)，其中 N 是数组中的元素数量。对于每一个元素 x，我们可以 O(1) 地寻找 target - x。
空间复杂度：O(N)，其中 N 是数组中的元素数量。主要为哈希表的开销。
*/
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
