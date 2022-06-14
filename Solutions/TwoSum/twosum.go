package TwoSum

// LeetCode - 1. Two Sum
// Time complexity == O(n²)

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[j] == target-nums[i] {
// 				r := []int{i, j}
// 				return r
// 			}
// 		}
// 	}
// 	return nil
// }

// Optimization - O(n)
// func twoSum(nums []int, target int) []int {
// 	ht := make(map[int]int)
// 	idxs := make([]int, 2)

// 	for i, j := range nums {
// 		diff := target - j

// 		if index, ok := ht[diff]; ok {
// 			idxs[0] = index
// 			idxs[1] = i
// 			break
// 		}

// 		ht[j] = i
// 	}
// 	return idxs
// }

// -- Fastest -- 
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}
