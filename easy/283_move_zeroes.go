package easy

// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements. Do this in-place
//
// Example:
// Input:  [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example:
// Input:  [0]
// Output: [0]
func moveZeroes(nums []int) {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
