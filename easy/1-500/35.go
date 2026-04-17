// https://leetcode.com/problems/search-insert-position/description/
package main

func searchInsert(nums []int, target int) int {
	leftI := 0
	rightI := len(nums)-1
    for {
		tempI := (leftI + rightI) / 2
		if nums[tempI] == target {
			return tempI
		} else if nums[tempI] < target {
			leftI = tempI + 1
		} else {
			rightI = tempI - 1
		}

		if leftI > rightI {
			return leftI
		}
	}
}
