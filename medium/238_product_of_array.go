package medium

// Given an integer array nums, return an array answer such that answer[i] is
// equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// Example: nums = [1,2,3,4]
// Output: [24,12,8,6]
//
// Example: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	left := 1
	for i := 0; i < length; i++ {
		result[i] = left
		left *= nums[i]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}
