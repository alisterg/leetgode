package easy

// Given two 0-indexed integer arrays nums1 and nums2, return a list answer of
// size 2 where:
// answer[0] is a list of all distinct integers in nums1 which are not present
// in nums2.
// answer[1] is a list of all distinct integers in nums2 which are
// not present in nums1.
func arrayDiff(a []int, b []int) [][]int {

	mapA := make(map[int]bool)
	mapB := make(map[int]bool)

	var resultA []int
	var resultB []int

	for _, v := range a {
		mapA[v] = true
	}
	for _, v := range b {
		mapB[v] = true
	}

	for _, v := range a {
		if !mapB[v] && mapA[v] {
			resultA = append(resultA, v)
			mapA[v] = false
		}
	}

	for _, v := range b {
		if !mapA[v] && mapB[v] {
			resultB = append(resultB, v)
			mapB[v] = false
		}
	}

	return [][]int{resultA, resultB}
}
