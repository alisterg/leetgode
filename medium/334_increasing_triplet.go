package medium

import "math"

// Given an integer array nums, return true if there exists a triple of
// indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
// If no such indices exists, return false.
//
// NOTE: it's not stated in the problem but the indexes do NOT have to
// be consecutive
//
// Example: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
//
// Example: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.
//
// Example: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
func increasingTriplet(nums []int) bool {
	smallest := math.MaxInt64
	secondSmallest := math.MaxInt64

	for _, num := range nums {
		if num <= smallest {
			smallest = num
		} else if num <= secondSmallest {
			secondSmallest = num
		} else {
			return true
		}
	}

	return false
}
