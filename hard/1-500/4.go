// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    sortedNums := []int{}

	nums1I := 0
	nums2I := 0

	m := len(nums1)
	n := len(nums2)

	for nums1I < m && nums2I < n {
		if nums1[nums1I] < nums2[nums2I] {
			sortedNums = append(sortedNums, nums1[nums1I])
			nums1I++
		} else {
			sortedNums = append(sortedNums, nums2[nums2I])
			nums2I++
		}
	}

	if nums1I < m {
		sortedNums = append(sortedNums, nums1[nums1I:]...)
	}
	if nums2I < n {
		sortedNums = append(sortedNums, nums2[nums2I:]...)
	}

	if len(sortedNums) % 2 == 0 {
		return float64(sortedNums[len(sortedNums)/2 - 1] + sortedNums[len(sortedNums)/2]) / 2
	} else {
		return float64(sortedNums[len(sortedNums)/2])
	}
}