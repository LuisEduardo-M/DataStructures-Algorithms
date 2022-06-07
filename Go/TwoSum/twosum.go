package TwoSum

// LeetCode - 1. Two Sum
// Time complexity == O(n²)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				r := []int{i, j}
				return r
			}
		}
	}
	return nil
}

// Optimization
