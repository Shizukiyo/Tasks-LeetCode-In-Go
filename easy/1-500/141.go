// https://leetcode.com/problems/linked-list-cycle/description/
package main

func hasCycle(head *ListNode) bool {
	oneStep := head
	twoStep := head

	for twoStep != nil && twoStep.Next != nil {

		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next

		if oneStep == twoStep {
			return true
		}
	}

	return false
}