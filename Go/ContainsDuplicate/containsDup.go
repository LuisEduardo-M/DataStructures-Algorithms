package ContainsDuplicate

// LeetCode - 217. Contains Duplicate
// Time Complexity - O(n)

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, contains := set[nums[i]]; contains {
			return true
		} else {
			set[nums[i]] = true
		}
	}
	return false
}
